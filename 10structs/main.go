package main

import "fmt"

func main() {
	// struct1()
	// struct2()
	struct3()
}

func struct1() {
	type User struct {
		Name    string
		Age     int
		Email   string
		IsAdmin bool
	}

	user := User{
		Name:    "Rick",
		Age:     23,
		Email:   "rick@gmail.com",
		IsAdmin: false,
	}

	fmt.Printf("Default struct: %v\n", user)
	fmt.Printf("With field names: %+v\n", user)
	fmt.Printf("Go syntax representation: %#v\n", user)
	
	var user1 User
	fmt.Printf("user1: %v\n", user1)
	fmt.Println(user1, user1.Name, user1.Age, user1.Email, user1.IsAdmin)

	userPtr := &user

	fmt.Println(userPtr, userPtr.Name)
	fmt.Println((*userPtr).Name)

}


func struct2()  {
	type Address struct {
		City    string
		Country string
	}

	type User struct {
		Name    string
		Age     int
		Address Address
		Skills  []string
		IsAdmin bool
		Metadata map[string]string
	}

	user := User{
		Name:    "Rick",
		Age:     23,
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

	fmt.Printf("%+v\n", user)

	// Access nested field
	fmt.Println(user.Address.City)

	fmt.Println(user.Metadata["role"])

	// Access slice field
	fmt.Println(user.Skills[1])

	// Add new field
	user.IsAdmin = true
	fmt.Println(user.IsAdmin)

	// Modify slice field
	user.Skills = append(user.Skills, "C++")
	fmt.Println(user.Skills)

	// Delete field
	user.Age = 0
	fmt.Println(user.Age)
}


func struct3()  {
	type User struct {
		Name    string
		Age     int
	}

	users := []User{
		{
			Name: "Rick",
			Age: 23,
		},
		{
			Name: "Alice",
			Age: 25,
		},
		{
			Name: "Bob",
			Age: 30,
		},
	}

	for _, user := range users {
		fmt.Println(user, user.Name, user.Age)
	}
}