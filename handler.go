package main

import (
	"encoding/base64"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var allParameters = map[string][2]string{
	"help":       {"Shows list of all commands.", "delfin help | delfin help <parameter>"},
	"version":    {"Shows delfin version and github url.", "delfin version"},
	"compress":   {"Compress a directory.", "delfin compress <directory location> <output location>"},
	"decompress": {"Decompress a .delfin file.", "delfin decompress <.delfin file location> <output location>"},
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
	case "compress":
		HandleCompress(parameters[1:])
		break
	case "decompress":
		HandleDecompress(parameters[1:])
		break
	case "version":
		HandleVersion()
		break
	default:
		fmt.Println("Parameter not found!")
		os.Exit(2)
	}
}

func HandleHelp(params []string) {
	if len(params) == 0 {
		fmt.Println("Please don't forget to check documentation on GitHub!\nList off all commands:")
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

func HandleCompress(params []string) {
	if len(params) < 2 {
		log.Fatal("Please provide input and output location.")
		os.Exit(2)
	} else {
		input, output := params[0], params[1]

		if CheckPath(input) == 1 && CheckPath(output) == 1 {
			var allFiles []string
			folder_splitted := strings.Split(input, "/")

			if len(folder_splitted) == 1 {
				folder_splitted = strings.Split(input, "\\")
			}

			file_name := folder_splitted[len(folder_splitted)-1]

			if file_name == "" {
				file_name = fmt.Sprintf("compressed_%s", time.Now().Format("2006-01-02_15-04-05"))
			}

			err := filepath.Walk(input, func(path string, info fs.FileInfo, err error) error {
				fmt.Println("Compressing", path)

				if CheckPath(path) == 2 {
					data, readError := os.ReadFile(path)

					if readError != nil {
						log.Fatal(fmt.Sprintf("Unexcepted Error on %s ", path), readError)
						os.Exit(4)
					}

					delfinData := DelfinData{
						EncodeByte(data),
						strings.Replace(path, input, file_name, -1),
						false,
					}

					allFiles = append(allFiles, delfinData.Format())
				} else if CheckPath(path) == 1 {
					delfinData := DelfinData{
						[]byte{},
						strings.Replace(path, input, file_name, -1),
						true,
					}

					allFiles = append(allFiles, delfinData.Format())
				}
				return err
			})

			if err != nil {
				log.Fatal("Unexcepted Error ", err)
				os.Exit(4)
			}

			data := EncodeByte([]byte(strings.Join(allFiles, "\n")))

			savedPath := fmt.Sprintf("%s/%s.delfin", output, file_name)

			err = os.WriteFile(savedPath, data, 0666)

			if err != nil {
				log.Fatal("Unexcepted Error While Saving the Compressed Folder ", err)
				os.Exit(4)
			}

			fmt.Println("Compressing is Completed. Please Check:", savedPath)
		} else {
			log.Fatal("Input or output location is not found or not a directory.")
			os.Exit(3)
		}
	}
}

func HandleDecompress(params []string) {
	if len(params) < 2 {
		log.Fatal("Please provide input and output location.")
		os.Exit(2)
	} else {
		input, output := params[0], params[1]

		if CheckPath(input) == 2 && CheckPath(output) == 1 {
			data, readError := os.ReadFile(input)

			if readError != nil {
				log.Fatal(fmt.Sprintf("Unexcepted Error on %s ", input), readError)
				os.Exit(4)
			}

			data = DecodeByte(data)
			var allDatas []DelfinData

			for _, line := range strings.Split(string(data), "\n") {
				splittedLine := strings.SplitN(line, ":", 3)

				convertedData, convertError := base64.StdEncoding.DecodeString(splittedLine[2])

				if convertError != nil {
					convertedData = []byte{}
				} else {
					convertedData = DecodeByte(convertedData)
				}

				var isDir bool

				if splittedLine[1] == "1" {
					isDir = true
				}

				allDatas = append(allDatas, DelfinData{
					convertedData,
					splittedLine[0],
					isDir,
				})
			}

			CreateDirs(allDatas, output)
			CreateFiles(allDatas, output)
			fmt.Println("Decompressing is Completed.")
		} else {
			log.Fatal("Input or output location is not found. / Input is not a file. / Output is not a directory.")
			os.Exit(3)
		}
	}
}

func HandleVersion() {
	fmt.Println("Delfin Version: 0.0.1\nGitHub: https://github.com/5elenay/delfin")
}
