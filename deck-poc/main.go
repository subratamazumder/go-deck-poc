package main

// import (
// 	"fmt"
// )


func main() {
	//d := newDeck()
	// d.print()
	// fmt.Println(d.toString())
	// hand, remainingDeck := deal(d,5)
	// hand.print()
	// remainingDeck.print()
	//d.saveToFile("deck.txt")
	cards := newDeckFromFile("deck.txt")
	cards.print() 
	cards.shuffle()
	cards.print()
}
