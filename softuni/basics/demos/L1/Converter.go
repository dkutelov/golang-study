package main

import "fmt"

func convert() {
	var inputCm float64
	fmt.Scanln(&inputCm)
	result := inputCm * 2.54
	fmt.Println(result)
}

