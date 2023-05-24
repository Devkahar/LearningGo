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

	var vegetableList [3]string
	vegetableList[0] = "Onion"
	vegetableList[1] = "Tomato"
	vegetableList[2] = "Ginger"

	fmt.Println("Fruit List ", vegetableList)
	fmt.Println("Fruit List length ", len(vegetableList))

}
