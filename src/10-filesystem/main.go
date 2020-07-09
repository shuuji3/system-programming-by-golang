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

// Append text to a file
func appendHello() {
	file, err := os.OpenFile("hello-go-file.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Appended Hello Go!!\n")
}

// Create directory
func createDir() {
	os.Mkdir("test", 0755)
	os.MkdirAll("test2/nest/dir/all", 0755)
}

func removeDir() {
	os.Remove("test")
	os.Remove("test2/nest/dir/all")
}

func move() {
	os.Rename("test2", "test3")
	os.Rename("hello-go-file.txt", "hello-go.txt")
}

func main() {
	open()
	read()
	appendHello()
	read()
	createDir()
	//removeDir()
	move()
}
