package main

import "fmt"
import "github.com/ian-antking/golang_curriculum/tickets_to_mars/tickets"

func main() {
	var generatedTickets []tickets.Ticket
	for count := 0; count <= 10; count += 1 {
		generatedTickets = append(generatedTickets, tickets.GenerateTicket())
	}

	for index := 0; index < len(generatedTickets); index += 1 {
		ticket := generatedTickets[index]
		ticketLine := fmt.Sprintf("%v  %v %v $ %v", ticket.Company, ticket.Days, ticket.Service, ticket.Price)
		fmt.Println(ticketLine)
	}
}