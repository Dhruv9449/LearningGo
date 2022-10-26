package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to math study!")

	// Sum
	var a int = 10
	var b int = 20

	fmt.Println("a + b =", a+b)

	// Random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// Random number using crypto/rand
	randomNo, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomNo)
}
