package main

import (
	"Learning_go/helper"
	"fmt"
	"strconv"
)

// syntatic sugar with : create assigning variable easier
const conferenceTickets uint = 50

/*
levels of scopes
 1. local: declaration within function, scope
 2. Package: can be used everywhere in the same package
 3. Global: Declaration outside all functions & uppercase first letter
    can be used across all package
*/
var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// since we are using slice, it can be dynamic
// hence, initialize with 0 initial size
var bookings = make([]map[string]string, 0)

// add a list/array to store booked user
// array should be given a fixed size and type
// var bookings []string
// alternative syntax
// var booking = []string{}
// booking := []string{}

func main() {
	/*
		func main is the entrypoint for our package.
	*/

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(remainingTickets, userTickets, email, firstName, lastName, conferenceName)

			firstNames := getFirstNames()
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

func greetUsers(confName string, confTickets uint, remainingTickets uint) {
	// Greeting message displayed in the start of application
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets with %v still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	// fmt.Printf("conferenceName is of type %T, conferenceTickrts is of type %T and remainingTickets of type %T\n", conferenceName, conferenceTickets, remainingTickets)

}

func getFirstNames() []string {
	// return first name of the entered user

	firstNames := []string{}
	// range can be used to iterate over different data structures
	// for array and slices, it return index and value
	for _, booking := range bookings {
		// fields function splits the string with white separator and return slice
		// fields comes from strings package
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, email string, firstName string, lastName string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for user
	// map uses key and value
	// map[key]value = this is only a type but we need to initialize with an empty map
	// make initalize value of given type
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	// Since we cannot mix match datatypes in map
	// we are converting the number into string
	// using the strconv package, we are converting int64 to string of base10
	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	// fmt.Printf("First value of Slice: %v\n", bookings[0])
	// fmt.Printf("Total value of Slice: %v\n", bookings)
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

}
