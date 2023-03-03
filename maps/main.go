package main

import "fmt"

func main() {

	marks := make(map[string]int)

	marks["English"] = 100
	marks["Math"] = 100

	printMap((marks))

}

func printMap(m map[string]int) {
	for mark, value := range m {
		fmt.Printf("Score for %s is %d \n", mark, value)
	}
}
