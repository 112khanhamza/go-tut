package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")

	// array {static}
	// var bookings [50] string

	// slice {dynamic}
	var bookings [] string

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

	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("The first type: %T\n", bookings)
	// fmt.Printf("The first type: %v\n", len(bookings))

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, userTickets, email)

	fmt.Printf("%v tickets remining for %v.\n", remainingTickets, conferenceName)

	fmt.Printf("These are all the booking in our application: %v\n", bookings)

}