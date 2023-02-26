package main

import "fmt"

var deckSize int

func main() {
	//var card string = "Ace of Spades";
	card := "Ace of Spades"
	card = "Five of Diamonds"
	//newCard := newCardFunc()
	loopTries()
	deckSize = 24
	fmt.Println(card)
	fmt.Println(deckSize)
	//fmt.Println(newCard)

}

// func newCardFunc() string {
// 	return "Five of Hearts"
// }

func loopTries() {
	//This is a slice as it is not having a fixed length
	//cards := []string{"Ace of Diamonds", newCardFunc()}
	//Use the deck type mentioned in the deck.go file
	//cards := deck{"Ace of Diamonds", newCardFunc()}
	//Cards append will return a new slice
	//cards = append(cards, "Six of Spades")

	//Now call the list of deck from the deck.go
	//cards := newd

	// for i, card := range cards {

	// 	fmt.Println(i, card)
	// }
	//print from the instance variable

	cards := newDeck()
	cards.print()

	//fmt.Println(cards)

}
