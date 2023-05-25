package main

import "fmt"

func main() {
	fmt.Println("Learning Loops and goto statement in GO")

	var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for _, day := range days {
		fmt.Println(day)
	}
	roughtValue := 1
	for roughtValue < 10 {
		if roughtValue == 2 {
			goto lco
		}
		if roughtValue == 5 {
			roughtValue++
			continue
		}
		fmt.Println("Value ", roughtValue)
		roughtValue++
	}
lco:
	fmt.Println("Jumping at dev@gm.in")
}
