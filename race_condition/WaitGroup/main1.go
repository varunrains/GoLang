package main1

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	//Message is accessed by two go routines which
	//will cause race condition
	msg = s
}

func main1() {
	msg = "Hello, World!"
	wg.Add(2)

	//go updateMessage("Hello, universe")
	//go updateMessage("Hello, cosmos!")

	wg.Wait()

	fmt.Println(msg)
}
