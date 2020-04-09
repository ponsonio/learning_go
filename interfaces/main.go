package main

import "fmt"

type englishBot struct{}

type spanishhBot struct{}

type bot interface {
	getGreeding() string
}

func (englishBot) getGreeding() string {
	return "hi there"
}

func (spanishhBot) getGreeding() string {
	return "hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeding())
}

func main() {

	eb := englishBot{}
	sb := spanishhBot{}

	printGreeting(eb)
	printGreeting(sb)

}
