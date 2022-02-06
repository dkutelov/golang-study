//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greeting(name string) {
	println("Hello,", name)
}
//* Write a function that returns any message, and call it from within
//  fmt.Println()
func getMessage() string {
	return "Any message"
}
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func add3Numbers(x, y, z float64) float64 {
	return x + y + z
}
//* Write a function that returns any number
func getRandomNumber(max int) int {
	return rand.Intn(max + 1)
}

//* Write a function that returns any two numbers
func getTwoRandomNumber(max int) (int, int) {
	n1, n2 := rand.Intn(max + 1), rand.Intn(max + 1)
	return n1, n2
}

func main() {
	greeting("Dary")

	fmt.Println(getMessage())

	result := add3Numbers(1.5,2.5,3.5)
	fmt.Println("The result is", result)

	fmt.Println("Random number",getRandomNumber(5))
	randNumber1, randNumber2 := getTwoRandomNumber(5)
	
	fmt.Println("Two random numbers are",randNumber1, "and",randNumber2)

	result2 := add3Numbers(float64(getRandomNumber(10)), float64(getRandomNumber(10)), float64(getRandomNumber(10)))
	fmt.Println("The sum of 3 random numbers is",result2)
}
