package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// https://www.youtube.com/watch?v=yyUHQIec83I&list=PLV1lYgSg14YwhmOYcpgJCiPjkvR-AdwRO&index=4

// package level variables
const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		// runs on a new thread
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
		}
	} else {
		if !isValidName {
			fmt.Println("First name and last name need to have at least 2 characters")
		}
		if !isValidEmail {
			fmt.Println("The email address provided does not contain @ character")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
		fmt.Println("Your input data is invalid, try again")
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Printf("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create a map for a user
	// all keys have the same type, all values have the same type
	// make() creates a new empty map
	//             type of key  type of value
	// var userData = make(map[string]string)

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Println("#### REMAINING ####")
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Println("#### ######### ####")
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func sendTicket(tickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", tickets, firstName, lastName)

	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("################")

	wg.Done()
}
