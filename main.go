package main

import (
	"log"

	"github.com/vitali-raikov/gowatest/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
