package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Creating web request in go")
	// performGetRequest()
	// performPostRequestWithJSON()
	performPostRequestWithFormData()
}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)

	// databytes, _ := ioutil.ReadAll(response.Body)

	// contentStr := string(databytes)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(contentStr)

}

func performPostRequestWithJSON() {
	myurl := "http://localhost:8000/post"

	payload := strings.NewReader(`
		{
			"name": "Dev",
			"age": 22
		}
	`)

	response, err := http.Post(myurl, "application/json", payload)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func performPostRequestWithFormData() {
	myurl := "http://localhost:8000/postform"

	payload := url.Values{}
	payload.Add("name", "Dev")
	payload.Add("age", "22")
	payload.Add("email", "devkahar@go.dev")

	response, err := http.PostForm(myurl, payload)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
