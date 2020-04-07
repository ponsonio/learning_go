package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "FF0000",
		"blue":  "HE4545",
		"white": "FFFFF",
	}

	printMap(colors)

	fmt.Println(colors)

	delete(colors, "blue")

	fmt.Println(colors)

	var co map[string]string

	fmt.Println(co)

	c := make(map[string]string)

	c["white"] = "JJJJJ"

	fmt.Println(c)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for ", color, " is ", hex)
	}
}
