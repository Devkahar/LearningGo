package main

import "fmt"

func main() {
	fmt.Println("Learning Arrays")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println("Fruit List ", fruitList)
	fmt.Println("Fruit List length ", len(fruitList))

	var vegetableList []string = []string{"Onion", "Tomato", "Ginger"}

	fmt.Println("Fruit List ", vegetableList)
	fmt.Println("Fruit List length ", len(vegetableList))

}
