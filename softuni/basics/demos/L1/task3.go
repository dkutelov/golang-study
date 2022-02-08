package main

import "fmt"

func taskThree() {
	var n1, n2 int

	fmt.Scan(&n1)
	fmt.Scan(&n2)

	sum := n1 * n2

	fmt.Println(sum)
}