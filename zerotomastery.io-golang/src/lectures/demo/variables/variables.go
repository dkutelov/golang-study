package main

import "fmt"

func main() {
	var myName = "Dari"
	fmt.Println("My name is", myName,".")

	var name string = "Katty"
	fmt.Println("My name is", name)

	username := "admin"
	fmt.Println("My name is", username)

	var sum int
	fmt.Println("The sum is", sum)

	//Comma ok
	part1, other := 1, 5
	fmt.Println("part 1 is", part1,"and other is", other)
	
	part2, other := 1, 0
	fmt.Println("part 2 is", part2,"and other is", other)

	// Block
	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("The lesson is about", lessonName,"of type", lessonType)

	// 'k' is lost
	word1, word2, _ := "hello", "there", '!'
	fmt.Println(word1, word2)
}
