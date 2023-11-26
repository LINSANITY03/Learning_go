package main

import "fmt"

func main() {
	/*
		func main is the entrypoint for our package.
	*/

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets with %v still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}
