package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/muya/coke/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
