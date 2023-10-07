package content

import (
	_ "embed"

	"github.com/gregoryv/web"
)

func NewContent() *Content {
	return &Content{}
}

type Content struct {
	title string
	cover *web.Element
	parts []*web.Element

	// user configurable
	style *web.CSS
}

func (c *Content) SetTitle(v string) {
	c.title = v
}

func (c *Content) SetCover(elements ...any) {
	c.cover = web.Wrap(elements...)
}

func (c *Content) Cover() *web.Element {
	return c.cover
}

func (c *Content) NewPart(elements ...any) {
	c.parts = append(c.parts, web.Wrap(elements...))
}

func (c *Content) Parts() []*web.Element {
	return c.parts
}

func (c *Content) SetStyle(v *web.CSS) {
	c.style = v
}

func (c *Content) Style() *web.CSS {
	return c.style
}
