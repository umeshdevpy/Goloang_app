package main

import "fmt"

func main(){
	conferenceName := "GoLangConference"
	totalNumberOfTickets := 50
	remainingNumberOfTickets := 50
	fmt.Printf("Welcome to %v booking site.\n", conferenceName)
	fmt.Printf("We have total of %v tickets. Available tickets are %v.\n", totalNumberOfTickets, remainingNumberOfTickets)
	fmt.Printf("conferenceName is %T. totalNumberOfTickets is %T. remainingNumberOfTickets is %T.\n", conferenceName, totalNumberOfTickets, remainingNumberOfTickets)

	var firstName string
	var lastName string
	var emailId string
	var numberOfTickets int

	fmt.Printf("Enter First Name: ")
	fmt.Scanln(&firstName)

	fmt.Printf("Enter Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Printf("Enter Email ID: ")
	fmt.Scanln(&emailId)

	fmt.Printf("Enter Number of Tickets: ")
	fmt.Scanln(&numberOfTickets)

	fmt.Printf("Hi %v %v. Welcome to %v. You have booked %v tickets. You will recive details to this %v shortly.\n", firstName, lastName, conferenceName, numberOfTickets, emailId)

}