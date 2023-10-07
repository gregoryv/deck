package view

import (
	_ "embed"
	"fmt"

	"github.com/gregoryv/web"
	. "github.com/gregoryv/web"
)

func NewPresentation() *Presentation {
	return &Presentation{
		css: onePageView().With(
			presentationView(),
			layoutView(),
		),
	}
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

var footerHeight int = 4  // vh
var headerHeight int = 12 // vh

func presentationView() *CSS {
	css := NewCSS()
	css.Style("html, body",
		"background-color: #fff",
	)

	css.Style(".page .view",
		"font-size: 3vh",
		"margin: 0 0",
		"padding: 0 0",
		//"padding: 0px 1.6vw 0px 1.6vw",
	)
	css.Style(".page .view .header",
		"text-align: center",
		"height: "+fmt.Sprintf("%vvh", headerHeight),
	)

	css.Style(".page .view .slide",
		"margin: auto",
		"padding: 1.6vw 1.6vw 1.6vw 1.6vw",
		"height: "+fmt.Sprintf("%vvh", 100-2*footerHeight-headerHeight-3),
		"overflow: hidden",
		//"border: 1px dashed red",
	)

	css.Style(".page .view .cover",
		"display: flex",
		"justify-content: center",
		"align-items: center",
		"height: "+fmt.Sprintf("%vvh", 100-2*footerHeight),
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
	css.Style(".page .view",
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
		//		"border: 1px solid red",
	)
	return css
}

type FooterFunc func(pageIndex int, parts []*Element) *Element

//go:embed enhance.js
var enhance []byte