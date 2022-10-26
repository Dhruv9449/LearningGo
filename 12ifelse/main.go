package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else study!")

	var result string
	var score = 80

	if score > 90 {
		result = "A"
	} else if score > 80 {
		result = "B"
	} else if score > 70 {
		result = "C"
	} else {
		result = "D"
	}

	fmt.Println(result)

	if num := 10; num > 0 {
		fmt.Println("num is positive")
	} else {
		fmt.Println("num is negative")
	}

}
