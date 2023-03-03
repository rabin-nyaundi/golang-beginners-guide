package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Todo struct {
	userId    int64
	id        int64
	title     string
	completed bool
}

type logWriter struct{}

func main() {

	// var todo Todo

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// byteSlice := make([]byte, 99999)

	// resp.Body.Read(byteSlice)

	// fmt.Println(string(byteSlice))
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return 1, nil
}
