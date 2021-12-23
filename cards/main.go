package main

import "fmt"

func main() {
	// initialization
	// var card string = "Five of Diamonds"
	card := newCard()

	// reassignment
	card = "Ace of Spades"

	fmt.Println(card)

	// slice
	cards := []string{newCard(), card, "Two of Hearts"}

	fmt.Println(cards)

	cards = append(cards, "Six of Spades")

	// loop index, value
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
