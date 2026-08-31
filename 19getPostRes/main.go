package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// performGetRequest()
	// performPostRequest()
	performPutRequest()
}

func performGetRequest() {
	// 1. Hit the URL
	const getURL = "https://jsonplaceholder.typicode.com/posts/1"

	response, err := http.Get(getURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 2. Inspect status code
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Type:", response.Header.Get("Content-Type"))

	// 3.1 Read the response body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Convert to String and Print
	bodyString := string(bodyBytes)
	fmt.Println("\nResponse Body:")
	fmt.Println(bodyString)

	// 3.2 Read body bytes using strings.Builder (efficient string allocation)
	var responseString strings.Builder

	byteCount, _ := responseString.Write(bodyBytes)
	fmt.Println("Byte Count:", byteCount)
	fmt.Println("Response Body:\n", responseString.String())
}

func performPostRequest() {
	const postURL = "https://jsonplaceholder.typicode.com/posts"

	// 1. Prepare raw JSON string payload
	requestBody := strings.NewReader(`
		{
			"title": "Learning Go Backend",
			"body": "Building robust web requests without Node.js",
			"userId": 1
		}
	`)

	// 2. Send POST request with content type
	response, err := http.Post(postURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 3. Check status code and read response
	fmt.Println("Status Code:", response.StatusCode)

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Created Record Response:\n", string(content))

}

func performPutRequest() {
	const formURL = "https://httpbin.org/post"

	// 1. Prepare form data using url.Values (map[string][]string)
	data := url.Values{}
	data.Add("firstname", "Atanu")
	data.Add("role", "Backend Developer")
	data.Add("language", "Golang")

	fmt.Println("data", data, "encoded data", data.Encode())

	// 2. Send the form request
	response, err := http.PostForm(formURL, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 3. Read the echoed response
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Form Post Response:\n", string(content))

}