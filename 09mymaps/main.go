package main

import "fmt"

func main() {
	fmt.Println("Now we are looking at maps/key-value pairs")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS is shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// To loop in maps

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
