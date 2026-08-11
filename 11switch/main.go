package main

import "fmt"

func main() {
	month := "January"
	switch month {
	case "January", "February":
		fmt.Println("Winter")
	case "March", "April", "May":
		fmt.Println("Spring")
	case "June", "July", "August":
		fmt.Println("Summer")
	case "September", "October", "November":
		fmt.Println("Autumn")
	default:
		fmt.Println("Invalid month")
	}

	age := 23

	switch {
	case age < 13:
		fmt.Println("Child")

	case age < 20:
		fmt.Println("Teenager")

	case age < 60:
		fmt.Println("Adult")

	default:
		fmt.Println("Senior")
	}
}
