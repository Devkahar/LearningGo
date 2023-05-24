package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learning Slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList: %T\n", fruitList)
	fmt.Println("FruitList ", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println("FruitList ", fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println("FruitList ", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 223
	highScores[1] = 999
	highScores[2] = 123
	highScores[3] = 150

	// highScores[4]=111;
	highScores = append(highScores, 555, 666, 777)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println("Sorted ", highScores)

	// how to remove a value form slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "nodejs", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
