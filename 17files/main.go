package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	fmt.Println("Working with files in Go")
	fileContent := "This is file content by dev"

	file, err := os.Create("./mylocgofile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, fileContent)
	checkNilErr(err)
	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile("./mylocgofile.txt")
}

func readFile(path string) {
	dataBytes, err := ioutil.ReadFile(path)
	checkNilErr(err)
	fmt.Println(string(dataBytes))
}
