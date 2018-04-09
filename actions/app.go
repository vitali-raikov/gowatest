package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/vitali-raikov/gowatest/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// T is the global translating variable
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_gowatest_session",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		var err error
		defaultLanguage := "en"

		if T, err = i18n.New(packr.NewBox("../locales"), defaultLanguage); err != nil {
			app.Stop(err)
		}
		app.Use(PersistLanguage)
		app.Use(T.Middleware())

		// Setup and use translations:
		app.Redirect(301, "/", fmt.Sprintf("/%s/customers", defaultLanguage))

		g := app.Group("/{language}")
		g.Resource("/customers", CustomersResource{&buffalo.BaseResource{}})

		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}

// PersistLanguage is handler to save language to session, as by default
// bufallo i18n reads from session, sort of hackish workaround
func PersistLanguage(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		s := c.Session()

		// Save language to session
		s.Set("lang", c.Param("language"))

		// Persist it in context variable across all views
		c.Set("currentLanguage", c.Session().Get(T.SessionName))

		err := s.Save()
		if err != nil {
			return err
		}

		return next(c)
	}
}
