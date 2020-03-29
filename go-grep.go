package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var searchword string = os.Args[1]
	var location string = os.Args[2]

	// Open location
	file, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//
	/* Get file names
	 files, err := cwd.Readdirnames(0)
	 if err != nil {
		log.Fatal(err)
	}

	// Open files successively
	for i := 0; i < len(files); i++ {
		current, err := os.Open(files[i])
		if err != nil {
			log.Fatal(err)
		}
	*/

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchword) {
			fmt.Printf("%v: %v\n", file.Name(), scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
}
