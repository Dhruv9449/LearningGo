package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study!")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	fmt.Println("The year is", presentTime.Year())
	fmt.Println("The month is", presentTime.Month())
	fmt.Println("The day is", presentTime.Day())

	createdDate := time.Date(2020, time.October, 21, 9, 13, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))
}
