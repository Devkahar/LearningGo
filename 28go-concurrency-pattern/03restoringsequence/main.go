package main

import (
	"fmt"
	"time"
)

type Message struct {
	Str  string
	Wait chan bool
}

func main() {
	// In Restoring Sequence we send channel on a channel
	// Make gorouting wait until its turn
	c := fanIn(getGreeting("Dev"), getGreeting("Ana"))

	for i := 0; i < 10; i++ {
		fmt.Println("In Main")
		msg1 := <-c
		fmt.Printf("Message1->%d: %s\n", i, msg1.Str)
		fmt.Println("----")
		msg2 := <-c
		fmt.Printf("Message2->%d: %s\n", i, msg2.Str)
		msg1.Wait <- true
		msg2.Wait <- true
	}
}

func fanIn(bob, ana <-chan Message) <-chan Message {
	ch := make(chan Message)

	go func() {
		for {
			ch <- <-bob
		}
	}()
	go func() {
		for {
			// Lets say ana take long time to reply
			time.Sleep(2100 * time.Millisecond)
			ch <- <-ana
		}
	}()
	return ch
}

func getGreeting(msg string) <-chan Message {

	ch := make(chan Message)
	blocker := make(chan bool)
	go func() {
		for {
			fmt.Println("In Greeting")
			ch <- Message{fmt.Sprintf("Hey %s, Good Morning", msg), blocker}
			time.Sleep(1 * time.Second)
			<-blocker
		}
	}()
	return ch
}
