package actions

import (
	"fmt"
	"time"

	"github.com/vitali-raikov/gowatest/models"
)

func (as *ActionSuite) Test_CustomersResource_List() {
	customer := as.CreateCustomer()

	res := as.HTML("/ru/customers").Get()
	as.Equal(200, res.Code)

	as.Contains(res.Body.String(), "Сменить язык")
	as.Contains(res.Body.String(), customer.FirstName)
	as.Contains(res.Body.String(), customer.LastName)
}

func (as *ActionSuite) Test_CustomersResource_Edit() {
	customer := as.CreateCustomer()

	res := as.HTML(fmt.Sprintf("/ru/customers/%s/edit", customer.ID)).Get()
	as.Equal(200, res.Code)

	as.Contains(res.Body.String(), customer.FirstName)
	as.Contains(res.Body.String(), customer.LastName)
}

func (as *ActionSuite) Test_CustomersResource_New() {
	res := as.HTML("/ru/customers/new").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_CustomersResource_Destroy() {
	customer := as.CreateCustomer()

	res := as.HTML("/ru/customers/%s", customer.ID).Delete()
	as.Equal(302, res.Code)
	as.Equal("/ru/customers", res.Location())
}

func (as *ActionSuite) Test_CustomersResource_Create() {
	customer := &models.Customer{
		EditPageDate: time.Now().UnixNano(),
		FirstName:    "Test",
		LastName:     "Testovich",
		Gender:       "male",
		BirthDate:    time.Now().AddDate(-20, 0, 0),
		Email:        "unexisting@email.com",
	}

	res := as.HTML("/ru/customers").Post(customer)
	as.Equal(302, res.Code)
	as.Equal("/ru/customers", res.Location())

	c := &models.Customer{}
	as.NoError(as.DB.First(c))
	as.Equal(customer.FirstName, c.FirstName)
}

func (as *ActionSuite) Test_CustomersResource_Create_Errors() {
	customer := &models.Customer{
		EditPageDate: time.Now().UnixNano(),
	}
	res := as.HTML("/ru/customers").Post(customer)
	as.Equal(422, res.Code)

	as.Contains(res.Body.String(), "First Name can not be blank.")
	as.Contains(res.Body.String(), "Last Name can not be blank.")
	as.Contains(res.Body.String(), "Email does not match the email format.")
	as.Contains(res.Body.String(), "Gender is not in the list [male, female].")
	as.Contains(res.Body.String(), "Birth Date must be before 60 years.")

	c, err := as.DB.Count(customer)
	as.NoError(err)
	as.Equal(0, c)
}
