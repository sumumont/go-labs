package main

import (
	"encoding/base64"
	"os"
)

func ReadImageBase64(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	str := base64.StdEncoding.EncodeToString(file)
	return str
}
