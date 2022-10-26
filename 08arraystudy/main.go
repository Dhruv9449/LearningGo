package main

import "fmt"

func main() {
	fmt.Println("Welcome to array study!")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Grapes"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"Potato", "Tomato", "Onion"}

	fmt.Println(vegList)
	fmt.Println(len(vegList))
}
