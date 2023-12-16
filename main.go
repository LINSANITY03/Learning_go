package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		func main is the entrypoint for our package.
	*/

	// syntatic sugar with : create assigning variable easier
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	bookings := []string{}

	// add a list/array to store booked user
	// array should be given a fixed size and type
	// var bookings []string
	// alternative syntax
	// var booking = []string{}
	// booking := []string{}

	fmt.Printf("conferenceName is of type %T, conferenceTickrts is of type %T and remainingTickets of type %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets with %v still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			// check the slice contents
			bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("First value of Slice: %v\n", bookings[0])
			// fmt.Printf("Total value of Slice: %v\n", bookings)
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			// range can be used to iterate over different data structures
			// for array and slices, it return index and value
			for _, booking := range bookings {
				// fields function splits the string with white separator and return slice
				// fields comes from strings package
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These first names of bookings are: %v\n", firstNames)
			if remainingTickets == 0 {
				// end program
				fmt.Println("Conference is booked out. Come back next year!.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			continue
		}
	}
}
