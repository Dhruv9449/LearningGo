package main

import "fmt"

func main() {
	fmt.Println(("Welcome to pointers!"))

	var ptr *int = new(int)
	fmt.Println(ptr)

	*ptr = 10
	fmt.Println("ptr =", ptr)

	num1 := 15

	var ptr2 *int = &num1

	fmt.Println(("Value of pointer is:"), ptr2)
	fmt.Println(("Value of pointer is:"), *ptr2)
}
