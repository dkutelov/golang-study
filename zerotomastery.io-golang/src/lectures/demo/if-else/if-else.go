package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 6

	if quiz1 > quiz2 {
		fmt.Println("1 is greater then 2")
	} else if quiz2 > quiz1 {
		fmt.Println("2 is greater then 1")
	} else {
		fmt.Println("Equal")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("More then 7")
	} else {
		fmt.Println("Less then 7")
	}
}
