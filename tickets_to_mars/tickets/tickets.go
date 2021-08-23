package tickets

type Ticket struct {
	Company string
	Days int
	Service string
	Price int
}

func Generate(companies []string, services []string, maxDuration int, maxPrice int, random func(int) int) Ticket {
	return Ticket{
		Company: companies[random(len(companies))],
		Days: random(maxDuration),
		Service: services[random(len(services))],
		Price: random(maxPrice),
	}
}