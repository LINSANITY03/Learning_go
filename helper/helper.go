package helper

import "strings"

// variable and function define outside any function can be accessed in all other files within the same package

// user capitalize function name to export expilictly on another file
// e.g. is Printf from fmt , we use as fmt.Printf
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
