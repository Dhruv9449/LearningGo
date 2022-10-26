package main

import "fmt"

func main() {
	fmt.Println("Functions in golang!")
	greeter()
	c := add(5, 6)
	fmt.Println(c)
	fmt.Println(proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func greeter() {
	fmt.Println("Hello")
}

func add(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (string, int) {
	total := 0
	for _, value := range values {
		total += value
	}
	return "Total:", total
}
