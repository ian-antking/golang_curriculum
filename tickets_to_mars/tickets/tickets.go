package tickets

import "fmt"

type Ticket struct {
	Company string
	Days int
	Service string
	Price int
}

func GenerateTicket(companies []string, services []string, maxDuration int, maxPrice int, random func(int) int) Ticket {
	return Ticket{
		Company: companies[random(len(companies) - 1)],
		Days: random(maxDuration),
		Service: services[random(len(services) - 1)],
		Price: random(maxPrice),
	}
}

func FormatTicket(ticket Ticket) string {
	return fmt.Sprintf("%v %v %v $ %v", ticket.Company, ticket.Days, ticket.Service, ticket.Price)
}