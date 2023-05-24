package main

import "fmt"

func main() {
	fmt.Println("Learning Maps in Go")

	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["PY"] = "python"
	languages["RB"] = "ruby"
	fmt.Println("List of languages ", languages)
	fmt.Println("JS full form ", languages["JS"])
	delete(languages, "RB")
	fmt.Println("List of languages ", languages)

	// Loop through languages
	for key, value := range languages {
		fmt.Printf("Key %v, Value %v\n", key, value)
	}
}
