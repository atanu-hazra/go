package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	// basicConversions()
	conversions()
}

func basicConversions() {

	// creating a reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Rate our app (1-5):")

	rating, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading rating:", err)
		return
	}

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println("Error parsing input to float:", err)
		return
	} else {
		fmt.Println("Added 1 to your rating", numRating + 1)
	}
}

func conversions() {

}


