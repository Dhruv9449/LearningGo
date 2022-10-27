package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://chitros.herokuapp.com/feed?limit=10&offset=20"

func main() {
	fmt.Println("Welcome to urls!")
	fmt.Println(URL)

	// Parse the URL
	result, _ := url.Parse(URL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	params := result.Query()
	fmt.Println(params)

	for param, value := range params {
		fmt.Println(param, value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "chitros.herokuapp.com",
		Path:    "/feed",
		RawPath: "user=Dhruv",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
