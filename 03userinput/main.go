package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// Coma ok || err ok
	input, _ := reader.ReadString('\n')

	fmt.Println("Hello", input)
	fmt.Printf("Type of input %T\n", input)
}
