package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("-----go-system-info-----")

	host, _ := os.Hostname()
	fmt.Println("Hostname:", host)
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)

	path, _ := os.LookupEnv("GOPATH")
	fmt.Println("GOPATH:", path)

	fmt.Println("Go version:", runtime.Version())
}
