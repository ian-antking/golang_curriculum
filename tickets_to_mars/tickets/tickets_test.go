package tickets

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTickets(t *testing.T) {
	t.Run("returns a Ticket struct", func(t *testing.T) {
		companies := []string{ "company0" }
		services  := []string{ "service0" }
		stubRandInt := func(number int) int {
			return number
		}

		ticket := Generate(companies, services, 1, 1, stubRandInt)

		assert.Equal(t, "company0", ticket.Company)
		assert.Equal(t, 1, ticket.Days)
		assert.Equal(t, "service0", ticket.Service)
		assert.Equal(t, 1, ticket.Price)
	})

	t.Run("can will select company and service from slices", func(t *testing.T) {
		companies := []string{ "company0", "company1", "company2" }
		services  := []string{ "service0", "service1", "service2" }
		stubRandInt := func(number int) int {
			return 1
		}

		ticket := Generate(companies, services, 1, 1, stubRandInt)

		assert.Equal(t, "company1", ticket.Company)
		assert.Equal(t, 1, ticket.Days)
		assert.Equal(t, "service1", ticket.Service)
		assert.Equal(t, 1, ticket.Price)
	})
}
