package tickets

import "math/rand"
import "time"

var companies = [3]string{"Virgin Galactic", "Space X", "Space Adventures"}
var services = [2]string{"round-trip", "one-way"}

type Ticket struct {
	Company string
	Days int
	Service string
	Price int
}

func GenerateTicket() Ticket {
	return Ticket{
		Company: companies[rand.Intn(len(companies))],
		Days: rand.Intn(365),
		Service: services[rand.Intn(len(services))],
		Price: rand.Intn(100),
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}