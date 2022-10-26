package main

import "fmt"

func main() {
	fmt.Println("Welcome to map study!")

	language := make(map[string]string)

	language["JS"] = "JavaScript"
	language["PY"] = "Python"
	language["GO"] = "GoLang"
	language["RB"] = "Ruby"

	fmt.Println("Language:", language)
	fmt.Println("Length:", len(language))

	fmt.Println("JS:", language["JS"])
	fmt.Println("PY:", language["PY"])

	delete(language, "JS")
	fmt.Println("Language:", language)

	for key, value := range language {
		fmt.Println(key, value)
	}
}
