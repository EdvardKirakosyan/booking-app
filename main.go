package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "GO conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	bookings := []string{}

	// fmt.Printf("conf tickets is %T, remaining tickets is %T, conf name Is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking app\n", conferenceName)
	fmt.Printf("We have %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets")

	for {
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

		// fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("The first type: %T\n", bookings)
		// fmt.Printf("The first value: %v\n", len(bookings))

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("thank you %v %v for booking %v tickets.\nYou will received them on %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("There are firstNames: %v\n", firstNames)
	}
}
