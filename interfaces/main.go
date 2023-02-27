package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	//Try to use one printGreeting with the help of interfaces
	//printGreeting(sb)

}

//Use interface
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//No need to have same method for
//almost same logic here the interfaces comes into picture.

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

//FUNCTION overloading is not supported in GO
//As it is supported in JAVA and C#
//Even though you have two method names with different set of
//parameters GO will through compiler error
//Check if method override is possible or not ?

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// One thing to notice for the receiver function
// Even if you dont use "eb" the compiler is not
// complaining about this
func (eb englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola amigo!"
}
