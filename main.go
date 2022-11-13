package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// array {static}
	// var bookings [50] string

	// slice {dynamic}
	var bookings [] string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		
		// ask user for their details
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want? ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		// array
		// bookings[0] = firstName + " " + lastName
		
		// slice
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, userTickets, email)

		fmt.Printf("%v tickets remining for %v.\n", remainingTickets, conferenceName)
		
		firstNames := [] string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These are all the first names of booking in our application: %v\n", firstNames)

		if remainingTickets == 0 {
			// end the program
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}

}