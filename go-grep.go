package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	var keyword string = os.Args[1]

	// Open CWD
	cwd, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}
	defer cwd.Close()

	// Get file names
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

		// Read file line by line
		scanner := bufio.NewScanner(current)
		for scanner.Scan() {
			// todo: Search for keyword, print if there is a success
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
