package tickets

import "math/rand"
import "time"

var companies = [3]string{"Virgin Galactic", "Space X", "Space Adventures"}
var services = [2]string{"round-trip", "one-way"}

type ticket struct {
	company string
	days int
	service string
	price int
}

func GenerateTicket() ticket {
	return ticket{
		company: companies[rand.Intn(len(companies))],
		days: rand.Intn(365),
		service: services[rand.Intn(len(services))],
		price: rand.Intn(100),
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}