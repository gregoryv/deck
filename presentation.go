package view

import (
	_ "embed"
	"fmt"

	"github.com/gregoryv/web"
	. "github.com/gregoryv/web"
)

func NewPresentation(scale int) *Presentation {
	// vh
	var footerHeight int = scale
	var headerHeight int = scale * 4
	var fontSize = scale

	css := NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
		"background-color: #fff",
	)
	css.Style(".page",
		"height: 100vh",
		"position: relative",
		"margin-bottom: 1vh",
	)
	css.Style(".page .view",
		"position: absolute",
		"top: 0",
		"left: 0",
		"right: 0",
		"bottom: "+vh(footerHeight),
	)
	css.Style(".page .footer",
		"position: absolute",
		"bottom: 0",
		"left: 0",
		"right: 0",
		"text-align: center",
		"height: "+vh(footerHeight),
		//		"border: 1px solid red",
	)
	css.Style(".page .view",
		"font-size: "+vh(fontSize),
		"margin: 0 0",
		"padding: 0 0",
		//"padding: 0px 1.6vw 0px 1.6vw",
	)
	css.Style(".page .view .header",
		"text-align: center",
		"height: "+vh(headerHeight),
	)
	css.Style(".page .view .slide",
		"margin: auto",
		"padding: 1.6vw 1.6vw 1.6vw 1.6vw",
		"height: "+vh(100-2*footerHeight-headerHeight-3),
		"overflow: hidden",
		//"border: 1px dashed red",
	)
	css.Style(".page .view .cover",
		"display: flex",
		"justify-content: center",
		"align-items: center",
		"height: "+vh(100-2*footerHeight),
		"text-align: center",
		//"border: 1px dashed red",
	)
	// toc
	css.Style(".toc a",
		"text-decoration: none",
	)
	css.Style(".h3",
		"margin-left: 5vw",
		"list-style-type: circle",
	)
	css = css.With(layoutView())
	return &Presentation{
		css: css,
	}
}

func vh(i int) string {
	return fmt.Sprintf("%vvh", i)
}

type Presentation struct {
	Title  string
	Author string

	cover *web.Element
	toc   *web.Element
	parts []*web.Element

	// user configurable
	css *CSS
}

func (p *Presentation) Style(x string, v ...string) {
	p.css.Style(x, v...)
}

func (p *Presentation) NewSlide(elements ...any) {
	if elements[0].(*Element).Name != "h2" {
		panic("h2 not first")
	}
	header := Div(Class("header"), elements[0])
	slide := Div(Class("slide"))
	slide.With(elements[1:]...)
	p.parts = append(p.parts, Wrap(header, slide))
}

func (p *Presentation) Document() *Page {
	body := Body()

	// create cover page if not set
	cover := p.cover
	if cover == nil {
		cover = Wrap(
			Div(Class("cover"),
				Table(
					Tr(
						Td(
							H1(p.Title),
						),
					),
					Tr(
						Td(
							p.Author,
						),
					),
				),
			),
		)
	}

	parts := p.parts
	// table of view
	toc := p.toc
	if toc == nil {
		ul := Ul()
		nav := Nav(
			Class("toc"),
			ul,
		)
		for i, root := range parts {
			WalkElements(root, func(e *Element) {
				if !(e.Name == "h2" || e.Name == "h3") {
					return
				}
				// +2 skip cover page and toc
				a := A(Href(fmt.Sprintf("#%v", i+3)), e.Text())
				ul.With(Li(Class(e.Name), a))
			})
		}

		toc = Wrap(
			Div(
				Class("header"),
				H2(p.Title),
			),
			Middle(30, nav),
		)
	}

	parts = append([]*Element{cover, toc}, parts...)
	for i, page := range parts {
		pageIndex := i + 1
		view := Div(Class("view"))
		view.With(page.Children...)

		body.With(
			Div(Class("page"), Attr("id", pageIndex),
				view,
				footer(pageIndex, parts),
			),
		)
	}

	return NewPage(
		Html(
			Head(
				Title(
					p.Title,
				),
				Style(p.css),
			),
			body,
			Script(
				string(enhance),
			),
		),
	)
}

func footer(pageIndex int, parts []*Element) *Element {
	return Div(Class("footer"),
		pageIndex, "/", len(parts),
	)
}

type FooterFunc func(pageIndex int, parts []*Element) *Element

//go:embed enhance.js
var enhance []byte
