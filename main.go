package main

import (
	"fmt"
	"strings"
)

// https://www.youtube.com/watch?v=yyUHQIec83I&list=PLV1lYgSg14YwhmOYcpgJCiPjkvR-AdwRO&index=4

func main() {

	// var conferenceName string = "Go Conference"
	// var conferenceName = "Go Conference"
	// Only for variables and does not allow to explicit assign a type
	conferenceName := "Go Conference"
	// Cannot change the value of constants
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	// var bookings = [50]string{"Nana", "Nicole"}
	// var bookings = [50]string{}
	// Array
	// var bookings [50]string
	// Slice -> dynamic list
	// var bookings []string
	// var bookings []string{}
	bookings := []string{}

	// fmt.Printf("conferenceTickets is %T, remaining tickets is %T, conferenceName is %T\n\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Welcome to the", conferenceName, "booking application")
	// fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// bookings[0] = "Nana"
	// bookings[10] = "Nicole"

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask the user for their name - need to provide a pointer
		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Printf("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Printf("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		// fmt.Println(remainingTickets)
		// fmt.Println(&remainingTickets)

		// userName = "Tom"
		// userTickets = 2
		// fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Slice type: %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		// fmt.Printf("These are all our bookings: %v\n", bookings)
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			// var firstName = names[0]
			// firstNames = append(firstNames, firstName)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

	}
}
