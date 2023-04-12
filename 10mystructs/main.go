package main

import "fmt"

func main() {
	fmt.Println("This is a class on structs")

	mayowa := User{"Mayowa", "mayowa@go.dev", true, 23}
	fmt.Println(mayowa)
	fmt.Printf("Mayowa's details are: %+v\n", mayowa)
	fmt.Printf("Name is %v and email is %v.", mayowa.Name, mayowa.Email)
}

// Note: No inheritance in golang. Also, no super or parent

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
