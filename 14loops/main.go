package main

import "fmt"

func main() {
	fmt.Println("For Loops in golang!")

	days := []string{"Monday", "Tuesday", "Wednesday", "Friday", "Sunday"}

	fmt.Println(days)

	// For loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println()

	// For loop using range
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println()

	// For each loop
	for _, day := range days {
		fmt.Println(day)
	}

	// While loop
	i := 0
	for i < 10 {
		if i == 7 {
			break
		}

		if i == 5 {
			goto here
		}

		if i == 3 {
			i++
			continue
		}

		fmt.Println(i)
		i++
	}

here:
	fmt.Println("Helllooo")
}
