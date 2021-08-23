package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
	"github.com/ian-antking/golang_curriculum/tickets_to_mars/tickets"
	"github.com/jedib0t/go-pretty/v6/table"
)

type configMap struct {
	Companies []string
	Services []string
	MaxDuration int
	MaxPrice int
}

var config configMap

func main() {
	rand.Seed(time.Now().UnixNano())
	configJson, _ := os.ReadFile("config.json")
	json.Unmarshal(configJson, &config)
	var generatedTickets []tickets.Ticket
	
	for count := 0; count <= 10; count += 1 {
		generatedTickets = append(
			generatedTickets, 
			tickets.Generate(
				config.Companies, 
				config.Services, 
				config.MaxDuration, 
				config.MaxPrice, 
				rand.Intn,
			),
		)
	}

	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)
	tableWriter.AppendHeader(table.Row{"Spaceline", "Days", "Service", "Price (million $)"})

	for _, ticket := range generatedTickets {
		tableWriter.AppendRow(table.Row{ticket.Company, ticket.Days, ticket.Service, ticket.Price})
	}

	tableWriter.Render()
}
