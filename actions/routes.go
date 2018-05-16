package actions

import "github.com/gobuffalo/buffalo"

func RoutesHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html", "old_application.html"))
}
