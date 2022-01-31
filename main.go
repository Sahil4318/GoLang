package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	//print(newDeckFromFile("my_cards"))
	// cards := newDeck()
	// fmt.Println(cards)
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	//cards.print()
	//hand, remainingCards := deal(cards, 5)
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	//hand.print()
	//remainingCards.print()
	// cards := deck{newCard(), newCard(), "Ace oF Diamonds"}
	// cards = append(cards, "Six of Spades") // does not modify original slice , assigns new slice to card
	// cards.print()
	//fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
