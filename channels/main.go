package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, url := range links {
		go checkLink(url, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(time.Second)
			checkLink(l, c)
		}(l)
	}

}

func checkLink(url string, c chan string) {
	_, error := http.Get(url)

	if error != nil {
		fmt.Println(url, " is down")
		c <- url
		return
	}
	fmt.Println(url, " is up")
	c <- url
}
