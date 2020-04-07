package main

func main() {

	//cards := newDeck()

	//hand, remaining := deal(cards, 5)

	//hand.print()
	//remaining.print()

	//cards.saveToFile("mycards.txt")
	cardfile := newDeckFromFile("mycards.txt")
	cardfile.shuffle()
	cardfile.print()
	//greeting := "hi there"
	//fmt.Println([]byte(greeting))
}
