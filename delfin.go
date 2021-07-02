package main

import (
	"bytes"
	"compress/zlib"
	"os"
)

func main() {
	HandleArguments()

	/* Compress a data with zlib:
	// Delfin will use this format

	var input bytes.Buffer
	data := []byte("test")
	writer := zlib.NewWriter(&input)
	writer.Write(data)
	writer.Close()

	ioutil.WriteFile("./test.txt", input.Bytes(), 0666)

	var output bytes.Buffer
	reader, _ := zlib.NewReader(&input)
	io.Copy(&output, reader)
	reader.Close()

	fmt.Println(out.String())

	ioutil.WriteFile("./text2.txt", out.Bytes(), 0666)
	*/
}

func CheckDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func EncodeByte(data []byte) []byte {
	var input bytes.Buffer

	writer := zlib.NewWriter(&input)
	writer.Write(data)
	writer.Close()

	return input.Bytes()
}
