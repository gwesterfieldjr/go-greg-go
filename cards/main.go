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
	// cards := []string{newCard(), card, "Two of Hearts"}
	cards := deck{newCard(), card, "Two of Hearts"}

	fmt.Println(cards)

	cards = append(cards, "Six of Spades")

	cards.print()

	deck := newDeck()
	deck.print()

	hand, remainingCards := deal(deck, 5)
	hand.print()
	remainingCards.print()

	fmt.Println(deck.toString())

	deck.saveToFile(("deck"))

	fmt.Println(newDeckFromFile("deck").toString())
}

func newCard() string {
	return "Five of Diamonds"
}
