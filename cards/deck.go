package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardValues {
		for _, value := range cardSuits {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards

}

// This is a receiver function can be called by the
// variables declared with
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Return multiple values from the function

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	//ioutil is a package which we have used to write the output to the file
	// 0666 is  a permission given to all for read/write
	//As "WriteFile" function return "error" so you will have to return the whole
	//statement to check the respective errors
	//ioutil.WriteFile is deprecated as of GO 1.16 so using os.WriteFile
	//save := ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
	//return save
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := os.ReadFile(fileName)
	if err != nil {
		//option 1: Log the error and return the new Deck of cards
		//Option 2: log the error and completely exit the program
		fmt.Println("Error:", err)
		//Value other than zero indicates non-success exit
		os.Exit(1)
	}

	stringSlice := strings.Split(string(byteSlice), ",")

	return deck(stringSlice)
}

func (d deck) shuffle() {
	length := len(d)

	//Other way of randomizing the values are as follows
	//Every time you run this program the seed should be different that means
	//you can use time to provide the initial seed value
	// source := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(source)
	// r.Intn(length - 1)

	for index := range d {
		randomIndex := rand.Intn(length - 1)
		d[index], d[randomIndex] = d[randomIndex], d[index]
	}
}
