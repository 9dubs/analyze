package cmd

import (
	"os"
	"fmt"
)

func getPath(name string) string {
	path:=wd + "\\" + name
	return path
}

func getFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	//defer file.Close()
	return file
}