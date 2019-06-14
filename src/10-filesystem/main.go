package main

import (
	"fmt"
	"io"
	"os"
)

// Create new file
func open() {
	file, err := os.Create("hello-go-file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Hello Go File!\n")
}

// Read file data
func read() {
	file, err := os.Open("hello-go-file.txt")
	if err != nil {
		panic(err)

	}
	defer file.Close()
	fmt.Println("Read data:")
	io.Copy(os.Stdout, file)
}

func main() {
	open()
	read()
}
