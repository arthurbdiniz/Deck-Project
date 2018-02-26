package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()
}
