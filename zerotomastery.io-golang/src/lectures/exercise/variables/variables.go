//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var color = "blue";
	fmt.Println("My favorite color is", color);
	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	birthYear, ageInYears := 1960, 52
	fmt.Println("My birth year is", birthYear,"and my age is", ageInYears)
	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial = "D"
		lastInitial = "K"
	)
	fmt.Println("My initials are", firstInitial, lastInitial)
	//* Declare (but don't assign!) a variable for your age (in days),
	var myAgeInDays int
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	myAgeInDays = ageInYears * 365;
	fmt.Println("My age in days is", myAgeInDays, "days")
	
}
