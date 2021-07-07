package main

import (
	"encoding/base64"
	"fmt"
)

type DelfinData struct {
	data        []byte
	path        string
	isDirectory bool
}

func (dd DelfinData) Format() string {
	isDir, data := 0, "-"

	if dd.isDirectory {
		isDir = 1
	}

	if len(dd.data) != 0 {
		data = base64.StdEncoding.EncodeToString(dd.data)
	}

	return fmt.Sprintf("%s:%d:%s", dd.path, isDir, data)
}

type Parameter struct {
	name, description, usage string
	function                 func(params []string)
}

type Metadata struct {
	Latest, Current, License string
}
