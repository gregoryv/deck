package deck

import (
	_ "embed"
	"fmt"
	"sync"

	. "github.com/gregoryv/web"
)

type Deck struct {
	Title  string
	Author string

	AutoCover bool
	cover     *Element

	AutoTOC bool
	toc     *Element

	cards []*Element

	mkuser sync.Once
	user   *CSS

	lastH2 *Element
}

func (p *Deck) CSS() *CSS {
	// vh
	var scale = 2
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
	)
	css.Style(".page .view",
		"font-size: "+vw(fontSize),
		"margin: 0 0",
		"padding: 0 0",
	)
	css.Style(".page .view .header",
		"text-align: center",
		"height: "+vh(headerHeight),
	)
	css.Style(".header .group",
		"font-size: 2vh",
	)
	css.Style(".header h3",
		"margin-top: 1vh",
	)
	css.Style(".page .view .slide",
		"margin: auto",
		"padding: 1.6vw 1.6vw 1.6vw 1.6vw",
		"height: "+vh(100-2*footerHeight-headerHeight-3),
		"overflow: hidden",
	)
	css.Style(".page .view .cover",
		"display: flex",
		"justify-content: center",
		"align-items: center",
		"height: "+vh(100-2*footerHeight-headerHeight),
		"text-align: center",
	)
	// toc
	css.Style(".toc",
		"position: absolute",
		"left: 20vw",
		"width: "+vw(55),
	)
	css.Style(".toc a",
		"text-decoration: none",
	)
	css.Style(".h3",
		"margin-left: 5vw",
		"list-style-type: circle",
	)
	css = css.With(layoutView())
	if p.user != nil {
		css = css.With(p.user)
	}
	return css
}

func (p *Deck) Style(x string, v ...string) {
	p.mkuser.Do(func() {
		p.user = NewCSS()
	})
	p.user.Style(x, v...)
}

func (p *Deck) NewCard(elements ...any) {
	c := p.newCard(elements...)
	p.cards = append(p.cards, c)
}

func (p *Deck) newCard(elements ...any) *Element {
	header := Div(Class("header"),
		p.headings(elements[0].(*Element)),
	)
	slide := Div(Class("slide"))
	slide.With(elements[1:]...)
	return Wrap(header, slide)
}

func (p *Deck) headings(e *Element) any {
	switch e.Name {
	case "h2":
		p.lastH2 = e
		return e

	case "h3":
		return Wrap(
			Span(Class("group"), p.lastH2.Children[0]),
			e,
		)
	default:
		return e
	}
}

func (p *Deck) Document() *Page {
	body := Body()
	var cards []*Element

	// create cover page if not set
	cover := p.cover
	if cover == nil && p.AutoCover {
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
	if cover != nil {
		cards = append(cards, cover)
	}

	// table of deck
	toc := p.toc
	if toc == nil && p.AutoTOC {
		ul := Ul()
		nav := Nav(
			ul,
		)
		for i, root := range cards {
			WalkElements(root, func(e *Element) {
				if !(e.Name == "h2" || e.Name == "h3") {
					return
				}
				// +2 skip cover page and toc
				a := A(Href(fmt.Sprintf("#%v", i+3)), e.Text())
				ul.With(Li(Class(e.Name), a))
			})
		}

		toc = p.newCard(
			H2(p.Title),
			Div(Class("toc"), nav),
		)
	}
	if toc != nil {
		cards = append(cards, toc)
	}
	cards = append(cards, p.cards...)

	for i, page := range cards {
		pageIndex := i + 1
		deck := Div(Class("view"))
		deck.With(page.Children...)

		body.With(
			Div(Class("page"), Attr("id", pageIndex),
				deck,
				footer(pageIndex, cards),
			),
		)
	}

	return NewPage(
		Html(
			Head(
				Title(
					p.Title,
				),
				Style(p.CSS()),
			),
			body,
			Script(
				string(enhance),
			),
		),
	)
}

func footer(pageIndex int, cards []*Element) *Element {
	return Div(Class("footer"),
		pageIndex, "/", len(cards),
	)
}

//go:embed enhance.js
var enhance []byte

func vh(i int) string {
	return fmt.Sprintf("%vvh", i)
}

func vw(i int) string {
	return fmt.Sprintf("%vvw", i)
}
