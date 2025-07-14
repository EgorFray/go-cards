package main


func main() {
	cards := newDeck()
	cards.print()

	cards.shuffle()
	cards.print()

	hand, left := deal(cards, 5)
	hand.print()
	left.print()

	hand.saveToFile("hand")
}