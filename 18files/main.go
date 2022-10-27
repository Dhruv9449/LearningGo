package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files!")
	content := "Heyy this is my first file"

	file, err := os.Create("./trial.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println(length)
	defer file.Close()

	readFile("./trial.txt")
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)

	checkNilErr(err)

	fmt.Println(data)
	fmt.Println(string(data))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
