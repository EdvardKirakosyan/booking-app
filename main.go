package main

import "fmt"

func main() {
	conferenceName := "GO conference"
	const conferenceTickets = 50
	remaininTickets := 50

	fmt.Printf("Welcome to our %v booking app\n", conferenceName)
	fmt.Printf("We have %v tickets and %v are still available\n", conferenceTickets, remaininTickets)
	fmt.Println("Get your tickets")

}
