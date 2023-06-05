package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition in go")
	// Race condtion occurs when two or more goroutines have shared data
	// and interact with it simultaneously.

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Add 1")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Add 2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Print @3")
		mut.RLock()
		fmt.Println("Printing @3 ", score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println(score)
}
