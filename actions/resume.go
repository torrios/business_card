package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/torrios/business_card/models"
	"github.com/pkg/errors"
)

func ResumeHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	var experience []models.Experience
	if err:= tx.All(&experience); err != nil {
		return errors.Wrap(err, "failed to load experience data")
	}
	c.Set("experience", experience)
	return c.Render(200, r.HTML("resume.html"))
}
