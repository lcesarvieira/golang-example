package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/brianvoe/gofakeit/v4"
)

// CustomersList default implementation.
func CustomersList(c buffalo.Context) error {

	// Create structs with random injected data
	type Customer struct {
		firstName string `fake:"{person.first}"`
		lastName  string `fake:"{person.last}"`
		address   string `fake:"{person.address}"`
	}

	var customers []Customer

	for i := 0; i < 10; i++ {
		customers = append( customers,Customer{
			firstName: gofakeit.FirstName(),
			lastName:  gofakeit.LastName(),
			address:   gofakeit.Address().Address,
		})	
	}
	
	return c.Render(200, r.JSON(customers))
}
