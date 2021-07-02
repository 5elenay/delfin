package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func main() {
	HandleArguments()
}

func CheckPath(path string) int {
	/*
		0 - Not Exists
		1 - Directory
		2 - File
	*/

	info, err := os.Stat(path)
	if err != nil {
		return 0
	}

	if info.IsDir() {
		return 1
	} else {
		return 2
	}
}

func EncodeByte(data []byte) []byte {
	var input bytes.Buffer

	writer := zlib.NewWriter(&input)
	writer.Write(data)
	writer.Close()

	return input.Bytes()
}

func DecodeByte(data []byte) []byte {
	input := *bytes.NewBuffer(data)
	var output bytes.Buffer

	reader, _ := zlib.NewReader(&input)
	io.Copy(&output, reader)
	reader.Close()

	return output.Bytes()
}

func CreateDirs(allDatas []DelfinData, output string) {
	for _, value := range allDatas {
		if value.isDirectory {
			fmt.Println("Create Directory:", value.path)
			err := os.MkdirAll(fmt.Sprintf("%s/%s", output, value.path), os.ModePerm)

			if err != nil {
				fmt.Println("Failed to Create Directory:", value.path, err)
			}
		}
	}
}

func CreateFiles(allDatas []DelfinData, output string) {
	for _, value := range allDatas {
		if !value.isDirectory {
			fmt.Println("Create File:", value.path)
			err := os.WriteFile(fmt.Sprintf("%s/%s", output, value.path), value.data, 0666)

			if err != nil {
				fmt.Println("Failed to Create File:", value.path, err)
			}
		}
	}
}
