package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice study!")

	var fruitList = []string{"Apple", "Orange", "Grapes"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println("Fruit List:", fruitList)

	fruitList = fruitList[1:3]
	fmt.Println("Fruit List:", fruitList)

	highScores := make([]int, 5)

	highScores[0] = 123
	highScores[1] = 534
	highScores[2] = 345
	highScores[3] = 456

	highScores = append(highScores, 101, 400)
	fmt.Println("High Scores:", highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)

	fmt.Println("High Scores:", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println()

	var courses = []string{"Python", "C++", "Java", "Go", "Ruby"}
	fmt.Println("Courses:", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses:", courses)

}
