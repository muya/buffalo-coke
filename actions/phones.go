package actions

import "github.com/gobuffalo/buffalo"

// PhonesShow default implementation.
func PhonesShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("phones/show.html"))
}

// PhonesIndex default implementation.
func PhonesIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("phones/index.html"))
}

