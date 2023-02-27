package main

import "fmt"

var deckSize int

func main() {
	//var card string = "Ace of Spades";
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	//newCard := newCardFunc()

	//loopTries()
	//conversionTest()
	//saveToFile()
	//readFromFile()
	//deckSize = 24
	// fmt.Println(card)
	// fmt.Println(deckSize)
	//fmt.Println(newCard)

	shuffle()
}

func shuffle() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//fmt.Println(cards.shuffle())
}

func readFromFile() {
	//This function will open the file and read from the file
	//First we need to convert from byte slice to string slice
	//and then convert it to deck type and call the reciever function "print" on cards
	cards := newDeckFromFile("myCards")
	cards.print()
}

// func newCardFunc() string {
// 	return "Five of Hearts"
// }

func saveToFile() {
	cards := newDeck()
	cards.saveToFile("myCards")
}

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
	// cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	//fmt.Println(cards)

}

func conversionTest() {
	cards := newDeck()
	fmt.Println(cards.toString())
}
