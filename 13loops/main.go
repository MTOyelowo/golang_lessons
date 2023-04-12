package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in GoLang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	fmt.Println("Loops example 1:")

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("Loops example 2:")

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("Loops example 3:")

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}

		if rogueValue == 5 {
			break
		}
		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeOnline")

	anotherValue := 1
	for anotherValue < 10 {
		if anotherValue == 5 {
			anotherValue++
			continue
		}
		fmt.Println("Value is: ", anotherValue)
		anotherValue++
	}
}
