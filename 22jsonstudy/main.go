package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON study")
	EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {
	Courses := []course{
		{"Go", 100, "Udemy", "1234", []string{"Go", "Golang"}},
		{"Python", 200, "Udemy", "1234", []string{"Python", "Django"}},
		{"JavaScript", 300, "Udemy", "1234", nil}}

	content, err := json.MarshalIndent(Courses, "", "  ")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(content))

}

func DecodeJSON() {
	jsonData := []byte(`
	{
		"name": "Go",
		"price": 100,
		"platform": "Udemy",
		"tags": ["Go", "Golang"]
	}
	`)

	var course course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &course)
		fmt.Println(course)
		fmt.Printf("%+v\n", course)
		fmt.Printf("Type: %T\n", course)
	} else {
		fmt.Println("JSON is invalid")
	}
}
