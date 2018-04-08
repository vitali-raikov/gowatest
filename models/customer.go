package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// Customer struct
type Customer struct {
	ID           uuid.UUID `json:"id" db:"id"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
	Email        string    `json:"email" db:"email"`
	Address      string    `json:"address" db:"address"`
	Gender       string    `json:"gender" db:"gender"`
	EditPageDate int64     `db:"-"`
	BirthDate    time.Time `json:"birth_date" db:"birth_date"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c Customer) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Customers is not required by pop and may be deleted
type Customers []Customer

// String is not required by pop and may be deleted
func (c Customers) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Customer) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.FirstName, Name: "First Name"},
		&validators.StringLengthInRange{Field: c.FirstName, Name: "First Name", Max: 100},

		&validators.StringIsPresent{Field: c.LastName, Name: "Last Name"},
		&validators.StringLengthInRange{Field: c.LastName, Name: "Last Name", Max: 100},

		&validators.StringLengthInRange{Field: c.Address, Name: "Address", Max: 200, Message: "Address must not be more than 200 characters"},

		&validators.TimeAfterTime{SecondTime: c.BirthDate, FirstTime: time.Now().AddDate(-18, 0, 0), FirstName: "Birth Date", SecondName: "18 years"},
		&validators.TimeIsBeforeTime{SecondTime: c.BirthDate, FirstTime: time.Now().AddDate(-60, 0, 0), FirstName: "Birth Date", SecondName: "60 years"},

		&validators.EmailIsPresent{Field: c.Email, Name: "Email"},
		&validators.StringInclusion{Field: c.Gender, Name: "Gender", List: []string{"Male", "Female"}},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Customer) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Customer) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
