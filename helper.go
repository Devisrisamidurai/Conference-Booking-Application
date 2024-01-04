package helper

import "strings"

func Validateuserinput(firstName string, lastName string, mail string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidmail := strings.Contains(mail, "@")
	isvalidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
	return isvalidName, isvalidmail, isvalidTicketNumber
}
