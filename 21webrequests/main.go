package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const BaseURL = "https://chitros.herokuapp.com/"

func main() {
	fmt.Println("Welcome to web requests")
	GetRequest(BaseURL)
	PostFormRequest(BaseURL + "signup")
}

func GetRequest(URL string, params ...string) {
	fmt.Println("Get request")

	response, err := http.Get(URL)
	CheckError(err)
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)
	CheckError(err)
	responseString.Write(content)
	// fmt.Println("Content: ", string(content))

	fmt.Println("Content :", responseString.String())

}

func PostFormRequest(URL string, params ...string) {
	fmt.Println("Post Form request")

	// Form data
	formData := url.Values{}
	formData.Add("username", "Test")
	formData.Add("fullname", "Test User")
	formData.Add("email", "test@example.com")
	formData.Add("password", "pass123")

	response, err := http.PostForm(URL, formData)
	CheckError(err)
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	content, _ := io.ReadAll(response.Body)
	fmt.Println("Content: ", string(content))
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
