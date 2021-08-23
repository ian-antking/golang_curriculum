package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/ian-antking/golang_curriculum/tickets_to_mars/tickets"
)

var companies []string = []string{"Virgin Galactic", "Space X", "Space Adventures"}
var services []string =  []string{"round-trip", "one-way"}

func main() {
	var generatedTickets []tickets.Ticket
	for count := 0; count <= 10; count += 1 {
		generatedTickets = append(generatedTickets, tickets.Generate(companies, services, 365, 100, rand.Intn))
	}

	for _, ticket := range generatedTickets {
		fmt.Println(tickets.Format(ticket))
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}