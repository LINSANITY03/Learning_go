package main

import "fmt"

func main() {
	/*
		func main is the entrypoint for our package.
	*/

	// syntatic sugar with : create assigning variable easier
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is of type %T, conferenceTickrts is of type %T and remainingTickets of type %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets with %v still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// can be used later

	// getting user input and saving the input in the firstName
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the amount of tickets: ")
	fmt.Scan(&userTickets)

	// getting pointer reference
	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
