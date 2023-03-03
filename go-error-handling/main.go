package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type File struct {
	name string
	data []byte
}

func main() {
	Open("test.txt")
}

func Open(name string) (file *File, err error) {
	f, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
		fmt.Println("error")
		return nil, err
	}

	defer f.Close()

	byteData := make([]byte, 1)
	// buf := new(bytes.Buffer)

	for {
		if _, err := f.Read(byteData); err != nil && errors.Is(err, io.EOF) {
			log.Fatal(err)
			fmt.Println("error")
			break
		}

		if err != nil {
			log.Fatal(err)
			fmt.Println("error")
			return nil, err

		}
	}
	// data, err := f.Read(byteData)

	fmt.Printf("%s\n", string(byteData))

	return file, nil
}
