package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	"os"
	"github.com/ian-antking/golang_curriculum/tickets_to_mars/tickets"
)

type configMap struct {
	Companies []string
	Services []string
	MaxDuration int
	MaxPrice int
}

var config configMap

func Init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	configJson, _ := os.ReadFile("config.json")
	json.Unmarshal(configJson, &config)
	var generatedTickets []tickets.Ticket
	for count := 0; count <= 10; count += 1 {
		generatedTickets = append(generatedTickets, tickets.Generate(config.Companies, config.Services, 365, 100, rand.Intn))
	}

	for _, ticket := range generatedTickets {
		fmt.Println(tickets.Format(ticket))
	}
}
