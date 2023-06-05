package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in go")

	// Create Channel
	ch := make(chan int, 2)

	// Wait groups of go routines

	wg := &sync.WaitGroup{}

	wg.Add(2)
	// recive only
	go func(myChan <-chan int, wg *sync.WaitGroup) {
		val, isChanelActive := <-ch
		fmt.Println(val, isChanelActive)
		wg.Done()
	}(ch, wg)
	// Send only
	go func(myChan chan<- int, wg *sync.WaitGroup) {
		myChan <- 5
		close(myChan)
		wg.Done()
	}(ch, wg)

	wg.Wait()

}
