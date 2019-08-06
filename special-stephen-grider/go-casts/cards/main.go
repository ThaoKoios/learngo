package main

func main() {
	//card := newCard()
	//fmt.Println(card)

	// cards := deck{"Ace of Diamonds", newCard()} // type deck []string
	// cards = append(cards, "Six of Spades")

	//cards := newDeckFromFile("my_cards")
	//cards.print()

	// cards.saveToFile("my_cards")
	//fmt.Println(cards.toString())
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards.print() // calling print() method on cards variable

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
