package main

import (
	"fmt"
)

func main() {
	sentence := "a bunch of words here"

	for index, letter := range sentence {
		if index%2 != 0 {
			fmt.Println("index :", index, "letter: ", string(letter))
		}
	}
}
