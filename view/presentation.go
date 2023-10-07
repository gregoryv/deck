package view

import (
	_ "embed"
	"fmt"

	"github.com/gregoryv/content"
	. "github.com/gregoryv/web"
)

func NewPresentation(c *content.Content) *Presentation {
	return &Presentation{c: c}
}

type Presentation struct {
	c *content.Content
}

func (p *Presentation) NewSlide(elements ...any) {
	if elements[0].(*Element).Name != "h2" {
		panic("h2 not first")
	}
	header := Div(Class("header"), elements[0])
	slide := Div(Class("slide"))
	slide.With(elements[1:]...)
	p.c.NewPart(header, slide)
}

func (p *Presentation) Document() *Page {
	c := p.c
	body := Body()
	parts := c.Parts()

	// create cover page if not set
	cover := c.Cover()
	if cover == nil {
		cover = Wrap(
			Div(Class("cover"),
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
				footer(c, pageIndex, parts),
			),
		)
	}

	return NewPage(
		Html(
			Head(
				Title(
					c.Title(),
				),
				Style(
					onePageView(),
					presentationView(),
					layoutView(),
					c.Style(),
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

var footerHeight int = 3  // vh
var headerHeight int = 12 // vh

func presentationView() *CSS {
	css := NewCSS()
	css.Style(".page .content",
		"font-size: 3vh",
		"margin: 0 0",
		"padding: 0 0",
		//"padding: 0px 1.6vw 0px 1.6vw",
	)
	css.Style(".page .content .header",
		"display: flex",
		"justify-content: center",
		"align-items: center",
		"height: "+fmt.Sprintf("%vvh", headerHeight),
	)

	css.Style(".page .content .slide",
		"margin: auto",
		"padding: 1.6vw 1.6vw 1.6vw 1.6vw",
		"height: "+fmt.Sprintf("%vvh", 100-2*footerHeight-headerHeight-3),
		"overflow: hide",
		//"border: 1px dashed red",
	)

	css.Style(".page .content .cover",
		"display: flex",
		"justify-content: center",
		"align-items: center",
		"height: "+fmt.Sprintf("%vvh", 100-2*footerHeight),
		"text-align: center",
		//"border: 1px dashed red",
	)
	return css
}

func onePageView() *CSS {
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
		"bottom: "+fmt.Sprintf("%vvh", footerHeight),
	)
	css.Style(".page .footer",
		"position: absolute",
		"bottom: 0",
		"left: 0",
		"right: 0",
		"text-align: center",
		"height: "+fmt.Sprintf("%vvh", footerHeight),
	)
	return css
}

type FooterFunc func(b *content.Content, pageIndex int, parts []*Element) *Element

//go:embed enhance.js
var enhance []byte
