package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/vitali-raikov/gowatest/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
