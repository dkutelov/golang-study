//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"math"
)
	

type Rectangle struct {
	x1,x2,x3,y1,y2,y3 float64
}

func getRectSides(rect Rectangle) (float64, float64) {
	sideA := math.Sqrt(math.Pow(rect.x2 - rect.x1, 2) + math.Pow(rect.y2 - rect.y1, 2))
	sideB := math.Sqrt(math.Pow(rect.x3 - rect.x2, 2) + math.Pow(rect.y3 - rect.y2, 2))

	return sideA, sideB
}

func getReactArea(rect Rectangle) float64 {
	sideA, sideB := getRectSides(rect)
	fmt.Println(sideA, sideB)
	return sideA * sideB
}

func getReactPerimeter(rect Rectangle) float64 {
	sideA, sideB := getRectSides(rect)
	return (sideA + sideB) * 2
}

func main() {
	rect1 := Rectangle {x1: 1, x2: 2, x3: 2, y1: 1, y2: 2, y3: 1}

	area := getReactArea(rect1)
	fmt.Println("The area of the reactangle is", area)

	perimeter := getReactPerimeter(rect1)
	fmt.Println("The perimeter of the reactangle is", perimeter)
}
