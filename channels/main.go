package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"https://jsonplaceholder.typicode.com/todos",
		"https://rickandmortyapi.com/api/character",
		"https://rickandmortyapi.com/api/location",
		"https://rickandmortyapi.com/api/episode",
		"https://restcountries.com/v3.1/all",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
		fmt.Println(<-c)
	}
}

func checkLink(linkUrl string, c chan string) {
	_, err := http.Get(linkUrl)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down\n", linkUrl)
		c <- "Link might be down"
		return
	}

	fmt.Printf("%s is up\n", linkUrl)
	c <- "Link is up"
}
