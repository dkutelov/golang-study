package main

import "fmt"


type Counter struct {
	hits int
}

// takes a pointer to a Counter
func increment(counter *Counter) {
	counter.hits += 1 // structures is placing *counter
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	// Function values are passed by value -> many compies
	// value := 10
	// var valuePtr *int
	// valuePtr = &value

	counter := Counter{}

	hello := "Hello"
	world := "World"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)
}
