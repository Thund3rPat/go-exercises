package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var pattern string
var filepath string
var caseinsensitive = flag.Bool("i", false, "Set this flag for case-insensitive-search")

// Parse Flags and arguments
func argparse() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("No arguments given!")
	} else if flag.Arg(1) == "" {
		log.Fatal("No File Path given")
	} else {
		pattern = flag.Arg(0)
		filepath = flag.Arg(1)
	}
}

// Search pattern in file
func searchfile(file *os.File) {
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

	// Open filepath
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	searchfile(file)
}
