package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/hectorrios/business_card/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
