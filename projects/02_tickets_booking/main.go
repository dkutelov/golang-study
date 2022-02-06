package main

import (
	"fmt"
	"strings"
)

// Package level variables need var and const, can not use :=
const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

// go run main.go
func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(firstName, lastName, email, userTickets)			
			fmt.Printf("The first names of bookings are: %v\n", getFirstNames())

			if (remainingTickets == 0) {
				fmt.Println("Our conference is sold out.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name are too short")
			}
			if !isValidEmail {
				fmt.Println("Email does not contain sign @")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets entered is invalid")
			}
		}
	}
}


func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint
	
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getFirstNames() []string {
	firstNames := []string {}

	for _, booking := range bookings {
		// Splits by space
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName, lastName, email string, userTickets uint) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidEmail, isValidName, isValidTicketNumber
}

func bookTicket(firstName, lastName, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}