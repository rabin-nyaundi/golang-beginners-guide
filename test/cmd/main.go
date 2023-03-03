package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	b1 := make([]byte, 1024*100)

	_, err = f.Read(b1)

	fmt.Println(string(b1))

}
