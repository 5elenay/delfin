package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type DelfinData struct {
	data         []byte
	path         string
	is_directory bool
}

var allParameters = map[string][2]string{
	"help":  {"Shows list of all commands.", "delfin help | delfin help <parameter>"},
	"zip":   {"Compress a directory.", "delfin zip <directory location> <output location>"},
	"unzip": {"Decompress a .delfin file.", "delfin unzip <.delfin file location> <output location>"},
}

func HandleArguments() {
	parameters := os.Args[1:]

	if len(parameters) == 0 {
		log.Fatal("You need to pass a parameter. Type 'help' for more information.")
		os.Exit(1)
	}

	parameter := strings.ToLower(parameters[0])

	switch parameter {
	case "help":
		HandleHelp(parameters[1:])
		break
	case "zip":
		HandleZip(parameters[1:])
		break
	default:
		fmt.Println(parameters)
	}
}

func HandleHelp(params []string) {
	if len(params) == 0 {
		fmt.Println("Please don't forget to check documentation at GitHub!\nList off all commands:")
		for key, value := range allParameters {
			fmt.Printf("    %s: %s\n", key, value[0])
		}
		fmt.Println("\nFor more information, Please type delfin help <parameter>.")
	} else {
		param := params[0]

		if value, status := allParameters[strings.ToLower(param)]; status {
			fmt.Printf("Description: %s\nUsage: %s\n", value[0], value[1])
		} else {
			log.Fatal("Parameter not found!")
			os.Exit(2)
		}
	}
}

func HandleZip(params []string) {
	if len(params) < 2 {
		log.Fatal("Please provide input and output location.")
		os.Exit(2)
	} else {
		input, output := params[0], params[1]

		if CheckDirectory(input) && CheckDirectory(output) {
			var all_files []string
			folder_splitted := strings.Split(output, "/")
			file_name := folder_splitted[len(folder_splitted)-1]

			if file_name == "" {
				file_name = fmt.Sprintf("compressed_%s", time.Now().Format("2006-01-02_15-04-05"))
			}

			fmt.Println(file_name)

			err := filepath.Walk(input, func(path string, info fs.FileInfo, err error) error {
				if CheckDirectory(path) == false {
					var input bytes.Buffer

					data, readError := ioutil.ReadFile(path)

					if readError != nil {
						log.Fatal("Unexcepted Error", readError)
						os.Exit(4)
					}

					writer := zlib.NewWriter(&input)
					writer.Write(data)
					writer.Close()

					formatted := fmt.Sprintf("%s:%d:%s", path, 0, input.String())
					all_files = append(all_files, formatted)
				} else {
					formatted := fmt.Sprintf("%s:%d:%s", path, 1, "-")
					all_files = append(all_files, formatted)
				}
				return err
			})

			if err != nil {
				log.Fatal("Unexcepted Error", err)
				os.Exit(4)
			}

			os.WriteFile(fmt.Sprintf("%s/%s.delfin", output, file_name), []byte(strings.Join(all_files, "\n")), 0666)
		} else {
			log.Fatal("Input or output location is not found / is not a directory.")
			os.Exit(3)
		}
	}
}
