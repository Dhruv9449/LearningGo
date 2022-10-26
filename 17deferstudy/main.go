package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer study!")

	fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Fourth")
	fmt.Println("Third")
	f1()
}

func f1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
