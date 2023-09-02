package main

// import "fmt"

func main() {
	cards := newDeckFromFile("my_card.txt")
	cards.print()
	cards.shuffle()
	cards.print()
	// hand, remainingCards := deal(cards, 5)
	// cards.saveToFile("my_card.txt")
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())
}
