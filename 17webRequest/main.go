package main

import (
	"fmt"
	"io"
	"net/http"
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

const url = "https://guthib.com/"

func main() {
	// 1. Send the GET request
	response, err := http.Get(url)
	checkNilErr(err)

	// 2. Important: Always close the connection stream when done
	defer response.Body.Close()

	// Without defer resp.Body.Close(): The server keeps holding memory, the OS runs out of socket descriptors, and your program eventually crashes under load.

	// 	A File Descriptor (FD) is a unique integer assigned by the operating system kernel to track an open I/O resource (files, standard input/output, pipes).

	// A Socket Descriptor is simply a specific type of file descriptor that represents an open network connection (a TCP/UDP endpoint with an IP and Port).

	// 3. Inspect the response metadata
	fmt.Printf("Response type: %T\n", response)
	fmt.Printf("Status code: %d\n", response.StatusCode)

	// 4. Read the response body stream into bytes
	dataBytes, err := io.ReadAll(response.Body)
	checkNilErr(err)

	// 5. Convert bytes to string to view the raw HTML/text
	content := string(dataBytes)
	fmt.Println("content: ", content)
}
