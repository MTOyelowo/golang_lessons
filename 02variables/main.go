package main

import "fmt"

const LoginToken string = "miaybb" // public variable: name starts with capital letter

func main() {
	var username string = "mayowa"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45567636783
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var smallFloat2 float64 = 255.45567636783
	fmt.Println(smallFloat2)
	fmt.Printf("Variable is of type: %T \n", smallFloat2)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherVariableStr string
	fmt.Println(anotherVariableStr)
	fmt.Printf("Variable is of type: %T \n", anotherVariableStr)

	// implicit type or manner of variable declaration

	var website = "spheresedjunior.com"
	fmt.Println(website)

	// no var style

	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
