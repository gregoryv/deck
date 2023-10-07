package book

import (
	_ "embed"

	. "github.com/gregoryv/web"
)

func NewBook() *Book {
	return &Book{}
}

type Book struct {
	pages []*Element
}

func (d *Book) Document() *Page {
	body := Body()
	for i, content := range d.pages {
		pageIndex := i + 1
		body.With(
			Div(Class("page"), Attr("id", pageIndex),
				content,

				Br(),
				Div(Class("footer"),
					pageIndex, "/", len(d.pages),
				),
			),
		)
	}

	p := NewPage(
		Html(
			Head(
				Style(theme()),
			),
			body,
			Script(
				string(enhance),
			),
		),
	)
	return p
}

func (d *Book) Page(elements ...any) {
	d.pages = append(d.pages, newPage(elements))
}

func newPage(elements []any) *Element {
	return Wrap(elements...)
}

// ----------------------------------------

func theme() *CSS {
	css := NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
	)
	css.Style(".page",
		"height: 100vh",
		"position: relative",
		debugBorder(),
	)
	css.Style(".page .footer",
		"position: absolute",
		"bottom: 0",
		"left: 0",
		"right: 0",
		"text-align: center",
	)
	return css
}

// ----------------------------------------

//go:embed enhance.js
var enhance []byte

var debug bool //= true

func debugBorder() (v string) {
	if debug {
		v = "border: 1px solid red"
	}
	return
}
