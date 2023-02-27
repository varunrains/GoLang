package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bs9234",
	}

	//second way of declaring maps
	//var colors map[string]string
	//One more way
	//colors := make(map[string]string)
	colors["yellow"] = "#ff0000"
	//To delete the key from the map
	delete(colors, "sdfs")
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}

}
