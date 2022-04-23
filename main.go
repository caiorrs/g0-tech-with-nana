package main

import (
	"fmt"
)

func main() {

	// var conferenceName string = "Go Conference"
	// var conferenceName = "Go Conference"
	// Only for variables and does not allow to explicit assign a type
	conferenceName := "Go Conference"
	// Cannot change the value of constants
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	// fmt.Printf("conferenceTickets is %T, remaining tickets is %T, conferenceName is %T\n\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Welcome to the", conferenceName, "booking application")
	// fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask the user for their name
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
}
