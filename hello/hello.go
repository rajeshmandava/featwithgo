package main

import (
	"fmt"
	"strings"
	//"os/user"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string

	for {
		var firstName string
		var lastName string
		var userTickets int
		var email string

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your second name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isUserTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isUserTicketNumber && isValidEmail {

			remainingTickets = remainingTickets - userTickets
			// bookings[5] = firstName + " " + lastName
			bookings = append(bookings, firstName+""+lastName)

			fmt.Printf("The whole array: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Array type: %T\n", bookings)
			fmt.Printf("Length of array: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			var noTicketsRemaining bool = remainingTickets > 0 || remainingTickets == 0
			if noTicketsRemaining {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @")
			}
			if !isUserTicketNumber {
				fmt.Println("Number of tickets you entered is invalid!")
			}
			fmt.Println("Your input data is invalid, try again!")
		}
	}

}
