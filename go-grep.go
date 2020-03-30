package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var pattern string
var fpath string
var caseinsensitive = flag.Bool("i", false, "Set this flag for case-insensitive-search")
var recursive = flag.Bool("r", false, "Set this flag for recursive-search")

// Outsource error handling
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Parse Flags and arguments
func argparse() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("No arguments given!")
	} else if flag.Arg(1) == "" {
		log.Fatal("No File Path given")
	} else {
		pattern = flag.Arg(0)
		fpath = flag.Arg(1)
	}
}

// Search pattern in file
func searchfile(f string) {
	file, err := os.Open(f)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentline := scanner.Text()

		if *caseinsensitive {
			currentline, pattern = strings.ToLower(currentline), strings.ToLower(pattern)
		}
		if strings.Contains(currentline, pattern) {
			fmt.Printf("%v: %v\n", file.Name(), scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error reading file: ", err)
	}
}

func main() {
	argparse()

	if *recursive {
		filelist := make([]string, 0)
		e := filepath.Walk(fpath, func(path string, info os.FileInfo, err error) error {
			check(err)
			if !info.IsDir() {
				filelist = append(filelist, path)
			}
			return nil
		})
		check(e)
		for _, v := range filelist {
			searchfile(v)
		}
	} else {
		searchfile(fpath)
	}
}
