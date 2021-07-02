package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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
			fmt.Println("Parameter not found!")
		}
	}
}
