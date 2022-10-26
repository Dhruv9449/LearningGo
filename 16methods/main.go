package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct study!")

	// No inheritance in golang.
	dhruv := User{"Dhruv", "Shah", 18, true}
	fmt.Println(dhruv)
	fmt.Printf("Details of %v is %+v\n", dhruv.FirstName, dhruv)
	fmt.Printf("Age of %v is %v\n", dhruv.FirstName, dhruv.GetAge())
	dhruv.NewAge()
	fmt.Printf("Age of %v is %v\n", dhruv.FirstName, dhruv.GetAge())

}

type User struct {
	FirstName string
	LastName  string
	Age       int
	Indian    bool
}

func (u User) GetAge() int {
	return u.Age
}

func (u User) NewAge() {
	u.Age = 20
	fmt.Println("New age is", u.Age)
}
