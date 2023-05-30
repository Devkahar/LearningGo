package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://loc.dev:3000/learn?coursename=reactjs&paymentid=123456"

func main() {
	fmt.Println("Learning how to handel urls in golang")

	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query parameters are: %T\n", qparams)
	fmt.Println(qparams)

	for key := range qparams {
		fmt.Println(qparams[key])
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "/tutcss",
		RawPath: "user=dev",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
