package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	//Message is accessed by two go routines which
	//will cause race condition+
	//this will be addressed by mutex - Mutual exclusion
	m.Lock() //Lock and unlock will be used to access the resources in a
	msg = s  //THread safe manner.
	m.Unlock()
}

func main() {
	msg = "Hello, World!"
	var mux sync.Mutex

	wg.Add(2)

	go updateMessage("Hello, universe1", &mux)
	go updateMessage("Hello, cosmos!", &mux)

	wg.Wait()

	fmt.Println(msg)
}
