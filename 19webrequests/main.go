package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://chitros.herokuapp.com/"

func main() {
	fmt.Println("Welcome to web requests!")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", response)
	fmt.Println(response)
	fmt.Println(response.Body)
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

}
