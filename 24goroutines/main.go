package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	fmt.Println("Learning Go routines")
	// go greetings("Hello")
	// greetings("Worlds")
	callWebpages()
}

// func greetings(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

var wg sync.WaitGroup // pointers
var mut sync.Mutex    // pointers
var signal []string

func callWebpages() {
	websites := []string{
		"https://amazon.in",
		"https://google.com",
		"https://go.dev",
		"https://lco.dev",
		"https://amazon.in",
	}
	// if we do not use go routines it take lot of time to execute code
	// but using goroutine we can create multiple threads and execute code faster
	for _, endpoint := range websites {
		go getStatusCode(endpoint)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signal)
}

func getStatusCode(endPoint string) {
	defer wg.Done()
	res, err := http.Get(endPoint)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d status code from %s\n", res.StatusCode, endPoint)
	mut.Lock()
	signal = append(signal, endPoint)
	mut.Unlock()
}
