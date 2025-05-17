package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var wd, err = os.Getwd()

func analyze(path string) {
	file := getFile(path)
	var linecounter int
	var wordcounter int
	linescanner := bufio.NewScanner(file)
	for linescanner.Scan() {
		linecounter++
		words := strings.Fields(linescanner.Text())
		wordcounter = wordcounter + len(words)
	}
	fmt.Println("Lines:", linecounter)
	fmt.Println("Words:", wordcounter)
}

func display(name string) {
	path := getPath(name)
	file, err := os.Stat(path)
	check(err)
	fmt.Println("Filename:", file.Name())
	fmt.Println("Path:", path)
	fmt.Println("Size:", file.Size(), "bytes")
	analyze(path)
	fmt.Println("Last Modified:", file.ModTime())
}

func renameFile(srcname string, dstname string) {
	src:= wd + "\\" + srcname
	dst:= wd + "\\" + dstname
	err := os.Rename(src, dst)
	check(err)
	fmt.Println("Renamed", srcname, "to", dstname)
}