package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://loc.dev"

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	fmt.Println("LOC Web Request")

	response, err := http.Get(url)

	checkNilErr(err)
	fmt.Printf("Response is of type %T\n", response)
	defer response.Body.Close() // callers responsibility to close connection

	databytes, err := ioutil.ReadAll(response.Body)

	checkNilErr(err)

	content := string(databytes)

	fmt.Println(content)
 
}
