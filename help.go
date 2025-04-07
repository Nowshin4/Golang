package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	bookings := make([]map[string]string, 0)

	greetUsers(conferenceName)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have %v tickets in total and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		var userName string
		var userEmail string
		var userPhone string
		var userTickets int
		var city string

		fmt.Println("Please enter your name")
		fmt.Scan(&userName)

		fmt.Println("Please enter your email")
		fmt.Scan(&userEmail)

		fmt.Println("Please enter your phone number")
		fmt.Scan(&userPhone)

		fmt.Println("Enter your ticket")
		fmt.Scan(&userTickets)

		fmt.Println("Please enter your city:")
		fmt.Scan(&city)

		// Validation checks
		isValidPhone := len(userPhone) > 5
		isValidEmail := strings.Contains(userEmail, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		isValidCity := city == "Singapore" || city == "London"

		if isValidPhone && isValidEmail && isValidTicketNumber && isValidCity {

			remainingTickets = remainingTickets - userTickets

			// Create a map for a user
			var userData = make(map[string]string)
			userData["userName"] = userName
			userData["userEmail"] = userEmail
			userData["userPhone"] = userPhone
			userData["city"] = city

			// Append the map of userData to bookings
			bookings = append(bookings, userData)

			fmt.Printf("Thank you %v for booking %v tickets, you will receive a confirmation email at %v soon\n", userName, userTickets, userEmail)
			fmt.Printf("%v tickets are left for %v\n", remainingTickets, conferenceName)
			fmt.Printf("The whole booking list: %v\n", bookings)

			if remainingTickets == 0 {
				fmt.Println("Our tickets are all booked, come next year")
				break
			}
		} else {
			fmt.Println("Your input data is invalid, try again")
		}
	}
}

func greetUsers(confName string) {
	fmt.Println("Hello there, welcome to Go Conference")
	fmt.Printf("This is our %v event\n", confName)
}
