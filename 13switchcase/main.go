package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in golang!")

	rand.Seed(time.Now().UnixNano())

	diceValue := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceValue)

	switch diceValue {
	case 1:
		fmt.Println("Dice value is 1, you can open!")
	case 2:
		fmt.Println("Dice value is 2, you can move 2 steps!")
	case 3:
		fmt.Println("Dice value is 3, you can move 3 steps!")
	case 4:
		fmt.Println("Dice value is 4, you can move 4 steps!")
	case 5:
		fmt.Println("Dice value is 5, you can move 5 steps!")
	case 6:
		fmt.Println("Dice value is 6, you can move 6 steps and roll the dice again!")
	default:
		fmt.Println("??")
	}
}
