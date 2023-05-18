package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza App")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate our app between 1 and 5")

	input, _ := reader.ReadString('\n')

	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("You rated , ", rating)
		fmt.Printf("Type of rating , %T\n", rating)
	}
}
