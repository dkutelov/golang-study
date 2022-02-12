package main

import (
	"fmt"
)

func add(lhs, rhs int) int {
	return lhs + rhs
}

type Operation func(lhs, rhs int) int

func compute(lhs, rhs int, op Operation) int {
	fmt.Printf("Running a computation with %v and %v.\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {
	fmt.Println("2 + 2 =", compute(2,2,add))

	fmt.Println("10 - 2 =", compute(10,2,func(lhs, rhs int) int {
		return lhs - rhs
	}))

	mul := func(lhs, rhs int) int {
		fmt.Printf("Multiplying %v * %v = ", lhs, rhs)
		return lhs * rhs
	}

	fmt.Println(compute(5,5,mul))
}
