package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to our app"
	fmt.Println(welcome)

	// creating a reader
	reader := bufio.NewReader(os.Stdin)

	// taking name input
	fmt.Println("Please enter your name:")

	nameInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	// taking age input
	fmt.Println("Please enter your age:")

	// ageInput, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading age:", err)
	// 	return
	// }

	// ageInput = strings.TrimSpace(ageInput)

	// age, err := strconv.Atoi(ageInput)
	// if err != nil {
	// 	fmt.Println("Invalid age:", err)
	// 	return
	// }

	
	age, err := ReadAge(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println("Thanks for using our app, ", nameInput, age)
}


func ReadAge(reader *bufio.Reader) (int, error) {

    input, err := reader.ReadString('\n')
    if err != nil {
        return 0, fmt.Errorf("reading age: %w", err)
    }

    input = strings.TrimSpace(input)

    age, err := strconv.Atoi(input)
    if err != nil {
        return 0, fmt.Errorf("converting age: %w", err)
    }

    return age, nil
}