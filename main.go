package main

import (
	"booking-app/helper"
	"fmt"
)

var conferenceName string = "GO conference"

const conferenceTickets int = 50

var remainingTickets int = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()

			fmt.Printf("There are firstNames: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("conf is over")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("fist & last you entered is to short")
			}

			if !isValidEmail {
				fmt.Println("email you entered is without @")
			}

			if !isValidTicketNumber {
				fmt.Println("number wrong")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking app\n", conferenceName)
	fmt.Printf("We have %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your firstName:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("list %v \n", bookings)

	fmt.Printf("thank you %v %v for booking %v tickets.\nYou will received them on %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
