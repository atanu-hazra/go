package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) Show() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Email:", u.Email)
	fmt.Println("Age:", u.Age)
}


func (u User) GetGreeting() {
	fmt.Println("Hello, my name is", u.Name, "and I am", u.Age, "years old")
}

// 1. Value Receiver: operates on a copy of the struct (cannot modify original)
func (u User) UpdateCopyEmail() {
	u.Email = "hello@new.mail" // this will not modify the original user
}

// 2. Pointer Receiver: operates on the original struct (can modify original)
func (u *User) IncreaseAge() {
	u.Age++
}

func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

func main() {
	user := User{
		Name:  "Alex",
		Email: "alex@example.com",
		Age:   25,
	}

	// Calling a value receiver method
	user.GetGreeting()

	user.UpdateCopyEmail()
	fmt.Printf("Updated User: %+v\n", user)

	// Calling pointer receiver methods (Go automatically takes the address &user)
	user.IncreaseAge()
	user.UpdateEmail("new@mail.com")

	fmt.Printf("Updated User: %+v\n", user)

	user.Show()
}
