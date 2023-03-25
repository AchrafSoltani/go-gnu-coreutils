package main

import (
	"fmt"
	"os"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// get Args
	argsWithoutProg := os.Args[1:]
	fmt.Print("cat ")
	fmt.Println(argsWithoutProg)
	file, err := os.Open("Makefile")
	defer file.Close()
	handleError(err)
	buffer5 := make([]byte, 5)
	n1, err := file.Read(buffer5)
	handleError(err)
	fmt.Printf("%d bytes: %s\n", n1, string(buffer5[:n1]))
}
