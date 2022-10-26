package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct study!")

	// No inheritance in golang.
	dhruv := User{"Dhruv", "Shah", 18, true}
	fmt.Println(dhruv)
	fmt.Printf("Details of %v is %+v\n", dhruv.FirstName, dhruv)
}

type User struct {
	FirstName string
	LastName  string
	Age       int
	Indian    bool
}
