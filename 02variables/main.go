package main

import "fmt"

var GloabalInt int = 43

func main() {
	var username string = "Nova"
	fmt.Println(username)
	fmt.Printf("Variable type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type: %T \n", isLoggedIn)

	var age int = 24
	fmt.Println(age)
	fmt.Printf("Variable type: %T \n", age)

	smallFloat := 334.4555
	fmt.Println(smallFloat)
	fmt.Printf("Variable type: %T \n", smallFloat)

	fmt.Println(GloabalInt)
	fmt.Printf("Variable type: %T \n", GloabalInt)
}
	