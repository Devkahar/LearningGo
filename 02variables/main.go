package main

import "fmt"

const LoginToken string = "xyz" // Public
func main() {
	var username string = "Dev Kahar"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T \n", smallVal)

	var smallFloat float32 = 255.490390490494430
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)

	var _64Float float64 = 255.490390490494430
	fmt.Println(_64Float)
	fmt.Printf("Variable is of type %T \n", _64Float)

	// Default values, for int 0, bool false, string "";

	var empty string
	fmt.Println(empty)
	fmt.Printf("Variable is of type %T \n", empty)

	// implicit type

	var guessType = "I gave string"
	fmt.Println(guessType)
	fmt.Printf("Variable is of type %T \n", guessType)

	// no var style

	name := "Dev Kahar"
	fmt.Println(name)
	fmt.Printf("Variable is of type %T \n", name)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T \n", LoginToken)

}
