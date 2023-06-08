package main

import (
	"fmt"
	"time"
)

func main() {
	ch := getGreeting("Dev")

	for {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("Greeting is slow")
			return
		}
	}
}

func getGreeting(msg string) <-chan string {

	ch := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("Hey %s, Good Morning", msg)
			time.Sleep(1 * time.Second)
		}
	}()
	return ch
}
