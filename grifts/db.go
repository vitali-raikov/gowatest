package grifts

import (
	"fmt"
	"strconv"
	"time"

	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
	"github.com/vitali-raikov/gowatest/models"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here

		for i := 0; i <= 42; i++ {
			customer := &models.Customer{
				EditPageDate: time.Now().UnixNano(),
				FirstName:    fmt.Sprintf("Test %s", strconv.Itoa(i)),
				LastName:     fmt.Sprintf("Testovich %s", strconv.Itoa(i)),
				Gender:       "male",
				BirthDate:    time.Now().AddDate(-20, 0, 0),
				Email:        fmt.Sprintf("unexisting%s@email.com", strconv.Itoa(i)),
			}

			if err := models.DB.Create(customer); err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	})

})
