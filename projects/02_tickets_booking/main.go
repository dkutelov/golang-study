package main

import (
	"fmt"
	"sync"
	"tickets_booking/helper"
	"time"
)

// Package level variables need var and const, can not use :=
const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
//var bookings = make([]map[string]string, 0) //List of maps with init size 0
var bookings = make([]User, 0) //List of structs with init size 0

// Create custom type
type User struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

// go run main.go
func main() {
	greetUsers()
	helper.PrintConferenceName(conferenceName)

	// Main treat does not wait for other and these are canceled
	// for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(firstName, lastName, email, userTickets)			
			//Go keyword creates a new treat
			wg.Add(1) //Add one treat that main treat should wait
			go sendTicket(firstName, lastName, email, userTickets)			
			fmt.Printf("The first names of bookings are: %v\n", getFirstNames())

			if (remainingTickets == 0) {
				fmt.Println("Our conference is sold out.")
				//break
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
	//}
	wg.Wait() // wait for all treats before you exit
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
		//var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func bookTicket(firstName, lastName, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	// Creates empty map with key string and value string
	//var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	
	var user = User {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	
	bookings = append(bookings, user)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// Run in separate threat
func sendTicket(firstName, lastName, email string, userTickets uint) {	
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("-----")
	fmt.Printf("Sending ticket:\n%v\nto email address%v\n", ticket, email)
	fmt.Println("-----")
	wg.Done() // removes the treat from the waiting list
}