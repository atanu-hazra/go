package main

import "fmt"

func main() {
	// struct1()
	// struct2()
	struct3()
}

// Demonstrates basic struct definition, instantiation, formatting verbs, and struct pointers
func struct1() {
	// Define a custom struct type
	type User struct {
		Name    string
		Age     int
		Email   string
		IsAdmin bool
	}

	// Instantiate struct using named fields
	user := User{
		Name:    "Rick",
		Age:     23,
		Email:   "rick@gmail.com",
		IsAdmin: false,
	}

	// Struct formatting verbs in fmt:
	fmt.Printf("Default struct: %v\n", user)             // output: Default struct: {Rick 23 rick@gmail.com false}
	fmt.Printf("With field names: %+v\n", user)          // output: With field names: {Name:Rick Age:23 Email:rick@gmail.com IsAdmin:false}
	fmt.Printf("Go syntax representation: %#v\n", user) // output: Go syntax representation: main.User{Name:"Rick", Age:23, Email:"rick@gmail.com", IsAdmin:false}

	// Zero-value struct: all fields get their respective type zero-values
	var user1 User
	fmt.Printf("user1: %v\n", user1)                                    // output: user1: { 0  false}
	fmt.Println(user1, user1.Name, user1.Age, user1.Email, user1.IsAdmin) // output: { 0  false}  0  false

	// Pointers to structs:
	// Go provides automatic dereferencing for struct pointers (userPtr.Name == (*userPtr).Name)
	userPtr := &user
	fmt.Println(userPtr, userPtr.Name) // output: &{Rick 23 rick@gmail.com false} Rick
	fmt.Println((*userPtr).Name)       // output: Rick
}

// Demonstrates nested structs, slices, maps within structs, and field modification
func struct2() {
	type Address struct {
		City    string
		Country string
	}

	type User struct {
		Name     string
		Age      int
		Address  Address           // Nested struct
		Skills   []string          // Slice inside struct
		IsAdmin  bool
		Metadata map[string]string // Map inside struct
	}

	user := User{
		Name: "Rick",
		Age:  23,
		Address: Address{
			City:    "New York",
			Country: "USA",
		},
		Skills: []string{"Golang", "Python", "Java"},
		Metadata: map[string]string{
			"role":   "developer",
			"status": "active",
		},
	}

	fmt.Printf("%+v\n", user) // output: {Name:Rick Age:23 Address:{City:New York Country:USA} Skills:[Golang Python Java] IsAdmin:false Metadata:map[role:developer status:active]}

	// 1. Access nested struct field
	fmt.Println(user.Address.City) // output: New York

	// 2. Access map field within struct
	fmt.Println(user.Metadata["role"]) // output: developer

	// 3. Access slice field within struct
	fmt.Println(user.Skills[1]) // output: Python

	// 4. Modify a boolean field
	user.IsAdmin = true
	fmt.Println(user.IsAdmin) // output: true

	// 5. Append to slice field inside struct
	user.Skills = append(user.Skills, "C++")
	fmt.Println(user.Skills) // output: [Golang Python Java C++]

	// 6. Reset/modify integer field
	user.Age = 0
	fmt.Println(user.Age) // output: 0
}

// Demonstrates a slice of structs and iterating over it
func struct3() {
	type User struct {
		Name string
		Age  int
	}

	// Slice of User structs
	users := []User{
		{
			Name: "Rick",
			Age:  23,
		},
		{
			Name: "Alice",
			Age:  25,
		},
		{
			Name: "Bob",
			Age:  30,
		},
	}

	// Iterating through the slice of structs
	for _, user := range users {
		fmt.Println(user, user.Name, user.Age)
		// output: {Rick 23} Rick 23 (1st iter)
		// output: {Alice 25} Alice 25 (2nd iter)
		// output: {Bob 30} Bob 30 (3rd iter)
	}
}