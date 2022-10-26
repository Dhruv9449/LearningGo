package main

import "fmt"

const LoginToken string = "1234"

func main() {
	var username string = "Dhruv"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T\n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFloat float32 = 3.1423423423423434
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	// Default values
	var defaultString string
	fmt.Println(defaultString)

	var defaultBool bool
	fmt.Println(defaultBool)

	var defaultInt int
	fmt.Println(defaultInt)

	var defaultFloat float64
	fmt.Println(defaultFloat)

	// Implicit type
	var implicitString = "Implicit String"
	fmt.Println(implicitString)

	// No var style
	// Does not work outside of a function
	noVarString := "No var string"
	fmt.Println(noVarString)

	// Constants
	fmt.Println(LoginToken)
}
