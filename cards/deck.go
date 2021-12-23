package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func newDeckFromFile(filePath string) deck {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to load deck from file: "+filePath, err)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filePath string) error {
	return ioutil.WriteFile(filePath, []byte(d.toString()), 0644)
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
