package main

import "fmt"

func main() {
	var conferenceName string = "GO conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conf tickets is %T, remaining tickets is %T, conf name Is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking app\n", conferenceName)
	fmt.Printf("We have %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets")

	var bookings = [50]string{}

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask for name
	fmt.Println("Enter your firstName:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - int(userTickets)

	fmt.Printf("thank you %v %v for booking %v tickets.\nYou will recieve them on %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for%v\n", remainingTickets, conferenceName)
}
