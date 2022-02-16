package utils

import (
	"fmt"
	"testing"
)

func TestFileObject(t *testing.T) {
	child := &FileObject{
		Subpath: "go-labs",
		Name:    "go-labs",
		IsDir:   true,
		Child:   nil,
	}
	err := child.ReadChild("/root/apulis/tmp")
	if err != nil {
		panic(err)
	}
	printFile(*child)
}
func printFile(fileObj FileObject) {
	fmt.Println(fileObj.Subpath, " ", fileObj.Name, " ", fileObj.IsDir)
	if len(fileObj.Child) != 0 {
		for _, file := range fileObj.Child {
			printFile(file)
		}
	}
}
