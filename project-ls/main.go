package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var showAll = flag.Bool("a", false, "Show all files")
	flag.Parse()

	for _, f := range listFiles("testdata", *showAll) {
		fmt.Println(f)
	}
}

func listFiles(dirname string, showAll bool) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !showAll && strings.HasPrefix(f.Name(), ".") {
			continue
		}
		dirs = append(dirs, f.Name())
	}

	return dirs
}
