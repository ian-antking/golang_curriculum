package tickets

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTickets(t *testing.T) {
	t.Run("returns a ticket struct", func(t *testing.T) {
		companies := []string{ "company1" }
		services  := []string{ "service1" }
		stubRandInt := func(number int) int {
			return number
		}
		ticket := GenerateTicket(companies, services, 1, 1, stubRandInt)

		assert.Equal(t, "company1", ticket.Company)
		assert.Equal(t, 1, ticket.Days)
		assert.Equal(t, "service1", ticket.Service)
		assert.Equal(t, 1, ticket.Price)
	})
}
