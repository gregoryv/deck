package view

import (
	_ "embed"

	"github.com/gregoryv/content"
	"github.com/gregoryv/web"
	. "github.com/gregoryv/web"
)

func NewPresentation(c *content.Content) *Presentation {
	return &Presentation{c: c}
}

type Presentation struct {
	c *content.Content
}

func (p *Presentation) Document() *Page {
	b := p.c
	c := p.c
	body := Body()
	parts := b.Parts()

	// create cover page if not set
	cover := b.Cover()
	if cover == nil {
		cover = Wrap(
			Div(Class("center"),
				Table(
					Tr(
						Td(
							H1(c.Title()),
						),
					),
					Tr(
						Td(
							c.Author(),
						),
					),
				),
			),
		)
	}

	parts = append([]*Element{cover}, parts...)
	for i, page := range parts {
		pageIndex := i + 1
		content := Div(Class("content"))
		content.With(page.Children...)
		body.With(
			Div(Class("page"), Attr("id", pageIndex),
				content,
				footer(b, pageIndex, parts),
			),
		)
	}

	return NewPage(
		Html(
			Head(
				Style(
					onePageView(),
					presentationView(),
					b.Style(),
				),
			),
			body,
			Script(
				string(enhance),
			),
		),
	)
}

func footer(b *content.Content, pageIndex int, parts []*Element) *Element {
	return Div(Class("footer"),
		pageIndex, "/", len(parts),
	)
}

func presentationView() *web.CSS {
	css := NewCSS()
	css.Style(".page .content",
		"font-size: 3vh",
		"padding: 0px 1.6vw 0px 1.6vw",
	)
	css.Style(".page .content .center",
		"display: flex",
		"justify-content: center",
		"align-items: center",
		"height: 80vh",
	)
	css.Style(".page .content .center table tr td",
		"text-align: center",
	)
	css.Style(".page .content h2",
		"text-align: center",
	)
	return css
}

func onePageView() *web.CSS {
	footerHeight := "25px"
	css := NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
	)
	css.Style(".page",
		"height: 100vh",
		"position: relative",
		"margin-bottom: 1vh",
	)
	css.Style(".page .content",
		//"padding: 10px 16px 10px 16px",
		"position: absolute",
		"top: 0",
		"left: 0",
		"right: 0",
		"bottom: "+footerHeight,
	)
	css.Style(".page .footer",
		"position: absolute",
		"bottom: 0",
		"left: 0",
		"right: 0",
		"text-align: center",
		"height: "+footerHeight,
	)
	return css
}

type FooterFunc func(b *content.Content, pageIndex int, parts []*Element) *Element

//go:embed enhance.js
var enhance []byte
