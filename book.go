package book

import (
	_ "embed"

	. "github.com/gregoryv/web"
)

func NewBook() *Book {
	return &Book{
		footer: defaultFooter,
	}
}

type Book struct {
	title     string
	coverPage *Element
	pages     []*Element

	// user configurable
	style  *CSS
	view   string
	footer FooterFunc
}

type FooterFunc func(b *Book, pageIndex int, pages []*Element) *Element

func (b *Book) SetTitle(v string) {
	b.title = v
}

func (b *Book) ViewAs(view string) {
	b.view = view
}

func (b *Book) SetStyle(v *CSS) {
	b.style = v
}

func (b *Book) Document() *Page {
	body := Body()
	pages := b.pages
	if b.coverPage != nil {
		pages = append([]*Element{b.coverPage}, pages...)
	}

	for i, page := range pages {
		pageIndex := i + 1
		content := Div(Class("content"))
		content.With(page.Children...)
		body.With(
			Div(Class("page"), Attr("id", pageIndex),
				content,
				b.footer(b, pageIndex, pages),
			),
		)
	}

	p := NewPage(
		Html(
			Head(
				Style(
					onePageView(),
					b.viewStyle(),
					b.userStyle(),
				),
			),
			body,
			Script(
				string(enhance),
			),
		),
	)
	return p
}

func defaultFooter(b *Book, pageIndex int, pages []*Element) *Element {
	return Div(Class("footer"),
		pageIndex, "/", len(pages),
	)
}

func (b *Book) CoverPage(elements ...any) {
	b.coverPage = Wrap(elements...)
}

func (b *Book) Page(elements ...any) {
	b.pages = append(b.pages, Wrap(elements...))
}

func (b *Book) userStyle() any {
	if b.style != nil {
		return b.style
	}
	return ""
}

func (b *Book) viewStyle() any {
	switch b.view {
	case "presentation":
		return presentationView()
	default:
		return ""
	}
}

// ----------------------------------------

func presentationView() *CSS {
	css := NewCSS()
	css.Style(".page .content",
		"font-size: 3vh",
	)
	return css
}

// ----------------------------------------

func onePageView() *CSS {
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

// ----------------------------------------

//go:embed enhance.js
var enhance []byte
