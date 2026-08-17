package main

import (
	"fmt"
	"io"
	"os"
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main()  {
	fileName := "example.txt"
	content := "Hello world"

	// 1. Create and Write to a File
	file, err := os.Create(fileName)
	checkNilErr(err)

	defer file.Close()

	// Write string content directly
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("file length: ", length)

	// Reading a file
	databyte, err := os.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Reading the file: ", string(databyte))
}
