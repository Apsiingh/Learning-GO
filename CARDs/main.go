package main

import "fmt"

func main() {

	// cards := newDeck()
	// saveToFile(cards, "my_cards.txt")
	// cards := newDeckFromFile("my_cards.txt")
	cards := newDeck()
	fmt.Println("Before shuffle")
	cards.printDeck()
	cards.shuffle()
	fmt.Println("=======================")
	fmt.Println("After shuffle")
	cards.printDeck()
}
