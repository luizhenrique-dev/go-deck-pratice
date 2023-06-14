package main

import "fmt"

// Iterando em na estrutura de dados: Slices
func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
}
