package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("------")
	cards.shuffler()
	cards.print()

}
