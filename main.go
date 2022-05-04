package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const confrenceTickets int = 50

var remainingTickets uint = 50
var confrenceName = "Go Confrence"
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var WG = sync.WaitGroup{}

func main() {

	greetUsers()

	// for {

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {

		bookedTicket(userTickets, firstName, lastName, email)
		WG.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// Call function print first Names
		firstNames := getFirstName()
		fmt.Printf("These first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("our Confrence is booked out. Come back next year\n")
			// break
		}
	} else {
		if !isValidName {
			fmt.Printf("first name or last name you entered is too short\n")
		}
		if !isValidEmail {
			fmt.Printf("Email address you entered doesn't contains @ sign\n")
		}
		if !isValidTickets {
			fmt.Printf("Number of tickets you entered is inValid\n")
		}
		fmt.Printf("Your Input data is Invalid\n")
	}
	// }
	WG.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confrenceTickets, remainingTickets)

}

func getFirstName() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first Name\n")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last Name\n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your Email\n")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets\n")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookedTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v.\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, confrenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(80 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("#####################################")
	fmt.Printf("Sending ticket %v \nto email address %v.\n", ticket, email)
	fmt.Println("#####################################")
	WG.Done()
}
