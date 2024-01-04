package main

import (
	"Booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const ConferenceTickets uint = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName   string
	lastName    string
	mail        string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, mail, userTicket := getuserinput()
	isvalidName, isvalidmail, isvalidTicketNumber := helper.Validateuserinput(firstName, lastName, mail, userTicket, remainingTickets)

	if isvalidName && isvalidmail && isvalidTicketNumber {

		bookticket(userTicket, firstName, lastName, mail)
		wg.Add(1)
		go sendTicket(userTicket, firstName, lastName, mail)

		firstNames := getfirstnames()
		fmt.Printf("The firstname of the bookings are:%v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("oops!All tickets got sold out! come back next year")
		}
	} else {
		if !isvalidName {
			fmt.Println("The entered name seems to be wrong .Please enter a valid name")
		}
		if !isvalidmail {
			fmt.Println("Please enter a valid mail")
		}
		if !isvalidTicketNumber {
			fmt.Println("Please enter a valid ticket number")
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("we have totally %v tickets and  %v tickets are still available\n", ConferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend\n")
	fmt.Println("*********************************")
}

func getfirstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}
func getuserinput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var mail string
	var userTicket uint

	fmt.Println("Enter your firstname:")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your lastname:")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scanln(&mail)

	fmt.Println("Enter no of tickets:")
	fmt.Scanln(&userTicket)

	return firstName, lastName, mail, userTicket
}
func bookticket(userTicket uint, firstName string, lastName string, mail string) {
	remainingTickets = remainingTickets - userTicket
	var userData = userData{
		firstName:   firstName,
		lastName:    lastName,
		mail:        mail,
		userTickets: userTicket,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thankyou %v%v for booking %v tickets for %v. You will recieve a confirmation mail at %v\n", firstName, lastName, userTicket, conferenceName, mail)
	fmt.Printf("%v tickets are available as of  now\n", remainingTickets)

}
func sendTicket(userTicket uint, firstName string, lastName string, mail string) {
	//send ticket to the user
	time.Sleep(30 * time.Second)
	var Ticket = fmt.Sprintf("%v tickets booked by %v%v", userTicket, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("sending ticket:\n %v \nto email address %v\n ", Ticket, mail)
	fmt.Println("#################")
	wg.Done()
}
