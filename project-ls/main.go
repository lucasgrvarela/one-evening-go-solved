package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	showHiddenFiles := flag.Bool("a", false, "show hidden files")
	flag.Parse()

	for _, file := range listFiles("testdata") {
		if strings.HasPrefix(file, ".") && !*showHiddenFiles {
			continue
		} else {
			fmt.Println(file)
		}
	}
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
