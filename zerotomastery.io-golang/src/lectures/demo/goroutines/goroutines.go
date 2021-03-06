package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	captIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		fmt.Printf("%c done!\n", r)
	}

	for i := 0; i < len(data); i++ {
		go captIt(data[i])
	}

	time.Sleep(100 * time.Microsecond)

	fmt.Printf("Capitalized: %c\n", capitalized)
	// result: Capitalized: [D B C A] or any other combination

	// run 15 times: 
	// for i in {1..15}; go run ./demo/goroutines
}
