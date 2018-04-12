package actions

import (
	"testing"
	"time"

	"github.com/gobuffalo/suite"
	"github.com/vitali-raikov/gowatest/models"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(App())}
	suite.Run(t, as)
}

func (as *ActionSuite) CreateCustomer() *models.Customer {
	customer := &models.Customer{
		FirstName: "Test",
		LastName:  "Testovich",
		Gender:    "male",
		BirthDate: time.Now().AddDate(-20, 0, 0),
		Email:     "unexisting@email.com",
	}

	verrs, err := as.DB.ValidateAndCreate(customer)
	as.NoError(err)
	as.False(verrs.HasAny())
	return customer
}
