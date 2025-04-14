package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck { 
	cards := deck{}
	cardSuits := []string{"spades", "diamonds", "hearts", "clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, " , ")
}

func saveToFile(d deck, filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option 1: log the error and exit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("File content in byte :", bs)
	deck := strings.Split(string(bs), " , ")
	return deck
}

func getRandomNo(d deck) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(len(d) - 1)
}

func (d deck) shuffle() {
	for i := range d {
		newPostion := getRandomNo(d)
		d[i], d[newPostion] = d[newPostion], d[i]
	}
}


