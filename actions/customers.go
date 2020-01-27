package actions

import (
	"github.com/gobuffalo/buffalo"
)

// CustomersList default implementation.
func CustomersList(c buffalo.Context) error {

	// Create structs with random injected data
	type Customer struct {
		firstName string `fake:"{person.first}"`
		lastName  string `fake:"{person.last}"`
		address   string `fake:"{person.address}"`
	}

	customers := Customer{}

	for i := 0; i < 50; i++ {
		// Pass your struct as a pointer
		//var f Customer{}
		customers = Customer{
			firstName: gofakeit.firstName(),
			lastName:  gofakeit.lastName(),
			address:   gofakeit.address(),
		}
	}

	//rest.Marshalled(c, &customers, http.StatusOK)
	return c.Render(200, r.JSON(&customers))
}
