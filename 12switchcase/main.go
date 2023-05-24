package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Cases in GO")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("You move to spot 1")
	case 2:
		fmt.Println("You move to spot 2")
	case 3:
		fmt.Println("You move to spot 3")
	case 4:
		fmt.Println("You move to spot 4")
	case 5:
		fmt.Println("You move to spot 5")
	case 6:
		fmt.Println("You move to spot 6")
	default:
		fmt.Println("Hey what was that!")
	}
}
