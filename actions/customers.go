package actions

import (
	"fmt"
	"time"

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
	return tx.Order("created_at desc")
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

	// We prefill this for search fields
	c.Set("lastName", lastName)
	c.Set("firstName", firstName)

	c.Set("customers", customers)
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("customers/index.html"))
}

// New is a hander for rendering new page
func (v CustomersResource) New(c buffalo.Context) error {
	customer := &models.Customer{}

	customer.BirthDate = time.Now().AddDate(-20, 0, 0)

	// Make customer available inside the html template
	c.Set("customer", customer)
	return c.Render(200, r.HTML("customers/new.html"))
}

// Create is a handler for processing customer creation
func (v CustomersResource) Create(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty customer
	customer := &models.Customer{}
	// Bind customer to the html form elements
	err := c.Bind(customer)
	if err != nil {
		return errors.WithStack(err)
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(customer)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make customer available inside the html template
		c.Set("customer", customer)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("customers/new.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Customer was created successfully")
	// and redirect to the customers index page
	return c.Redirect(302, fmt.Sprintf("/%s/customers", c.Session().Get(T.SessionName)))
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

// Destroy is a hander for rendering edit page
func (v CustomersResource) Destroy(c buffalo.Context) error {
	customer := &models.Customer{}

	err := v.scope(c).Find(customer, c.Param("customer_id"))
	if err != nil {
		return c.Error(404, err)
	}

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	err = tx.Destroy(customer)
	if err != nil {
		return errors.WithStack(err)
	}
	// If there are no errors set a flash message
	c.Flash().Add("success", T.Translate(c, "alert_customers_destroyed"))
	// Redirect to the customers index page
	return c.Redirect(302, fmt.Sprintf("/%s/customers", c.Session().Get(T.SessionName)))
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

	if customer.EditPageDate != customer.UpdatedAt.UnixNano() {
		c.Set("customer", customer)
		c.Flash().Add("warning", T.Translate(c, "alert_customers_changed"))
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
	c.Flash().Add("success", T.Translate(c, "alert_customers_save"))
	// and redirect to the customers index page

	return c.Redirect(302, fmt.Sprintf("/%s/customers", c.Session().Get(T.SessionName)))
}
