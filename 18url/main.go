package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://guthib.com:3000/learn?coursename=reactjs&paymentid=ghbj454"

func main() {
	// 1. Parsing the URL string
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}

	// 2. Extracting core components
	fmt.Println("Scheme:", result.Scheme)       // https
	fmt.Println("Host:", result.Host)           // guthib.com:3000
	fmt.Println("Hostname:", result.Hostname()) // guthib.com
	fmt.Println("Port:", result.Port())         // 3000
	fmt.Println("Path:", result.Path)           // /learn
	fmt.Println("RawQuery:", result.RawQuery)   // coursename=reactjs&paymentid=ghbj454

	// 3. Extracting and iterating over Query Parameters
	queryParams := result.Query() // type: url.Values
	fmt.Printf("Query Params Type: %T\n", queryParams)

	// Access specific query values
	fmt.Println("Course Name:", queryParams["coursename"][0])
	fmt.Println("Payment ID:", queryParams.Get("paymentid"))

	// Iterate through all query params
	for key, val := range queryParams {
		fmt.Printf("Param '%s': %v\n", key, val)
	}

	// 4. Constructing a URL from scratch
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "guthib.com",
		Path:     "/tutorials",
		RawQuery: "user=nova",
	}

	constructedUrl := partsOfUrl.String()
	fmt.Println("\nConstructed URL:", constructedUrl)
}