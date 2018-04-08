package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/pkg/errors"
	"github.com/vitali-raikov/gowatest/models"
)

// CustomersResource is the resource for the customers model
type CustomersResource struct {
	buffalo.Resource
}

func (v CustomersResource) scope(c buffalo.Context) *pop.Query {
	tx := c.Value("tx").(*pop.Connection)
	return tx.Order("id desc")
}

// List is a default handler to serve home page
func (v CustomersResource) List(c buffalo.Context) error {
	customers := []models.Customer{}

	firstName := c.Param("first_name")
	lastName := c.Param("last_name")

	q := v.scope(c)

	if firstName != "" {
		q = q.Where("lower(first_name) LIKE lower(?)", fmt.Sprintf("%%%s%%", firstName))
	}

	if lastName != "" {
		q = q.Where("lower(last_name) LIKE lower(?)", fmt.Sprintf("%%%s%%", lastName))
	}

	q = q.PaginateFromParams(c.Params())

	err := q.All(&customers)
	if err != nil {
		return errors.WithStack(err)
	}

	c.Set("lastName", lastName)
	c.Set("firstName", firstName)

	c.Set("customers", customers)
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("customers/index.html"))
}

// Edit is a hander for rendering edit page
func (v CustomersResource) Edit(c buffalo.Context) error {
	customer := &models.Customer{}
	err := v.scope(c).Find(customer, c.Param("customer_id"))
	if err != nil {
		return c.Error(404, err)
	}
	// Make customer available inside the html template
	c.Set("customer", customer)
	return c.Render(200, r.HTML("customers/edit.html"))
}

// Update is handler for processing customer updating
func (v CustomersResource) Update(c buffalo.Context) error {
	customer := &models.Customer{}
	err := v.scope(c).Find(customer, c.Param("customer_id"))

	if err != nil {
		return c.Error(404, err)
	}

	// Bind Customer to the html form elements
	err = c.Bind(customer)
	if err != nil {
		return errors.WithStack(err)
	}

	fmt.Println(customer.EditPageDate)
	fmt.Println(customer.UpdatedAt.UnixNano())
	if customer.EditPageDate != customer.UpdatedAt.UnixNano() {
		c.Set("customer", customer)
		c.Flash().Add("warning", "Customer was updated meanwhile, please confirm that you still want to update customer")
		return c.Render(422, r.HTML("customers/edit.html"))
	}

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	verrs, err := tx.ValidateAndUpdate(customer)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make customer available inside the html template
		c.Set("customer", customer)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("customers/edit.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Customer was updated successfully")
	// and redirect to the customers index page
	return c.Redirect(302, "/ru/customers")
}
