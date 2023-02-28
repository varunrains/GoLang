package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//The order of execution is the order in which the list is created!
	//There will be delay in getting the address (sequence)
	//This is where go routine will come handy
	//Channel is the only way to communicate between the go routines
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://go.dev",
		"http://amazon.com",
	}

	//You are creating a channel of type string
	//Note you cannot pass different data type to this channel
	channel := make(chan string)

	for _, link := range links {
		//You can only use go keyword in front of the function
		//go checkLink(link, channel)
		go checkMultipleLink(link, channel)

	}

	// for i := 0; i < len(links); i++ {
	// 	//The below call is the blocking call
	// 	//The main routine will wait until it receives messages from the channel
	// 	//fmt.Println(<-channel)

	// 	go checkMultipleLink(<-channel, channel)
	// }

	//Infinite loop in GO with for
	// for {
	// 	//The below call is the blocking call
	// 	//The main routine will wait until it receives messages from the channel
	// 	//fmt.Println(<-channel)

	// 	go checkMultipleLink(<-channel, channel)
	// }

	//To write a clean infinite  loop with channels:
	for l := range channel {
		//Sleep the go routine as you dont want to hit these servers multiple times
		//The go routine will sleep for 5 seconds
		//This is a bad practice to put the sleep in the main go routine
		//It will pause all the messages that we receive through the channel
		//time.Sleep(time.Second * 5)
		//Use a functional literal "lamda" function like in C#
		//go checkMultipleLink(l, channel)
		//If you dont pass the argument then the l will be same
		//across the go routines and it will lead into errors
		//Even the compile will give the errors when you do that
		go func(cd string) {
			time.Sleep(time.Second * 5)
			checkMultipleLink(cd, channel)
		}(l)
	}

}

func checkMultipleLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

func checkLink(link string, c chan string) {
	//sleep also here is not a  good approach
	//time.Sleep(time.Second * 5)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down!"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Site is up!"
}
