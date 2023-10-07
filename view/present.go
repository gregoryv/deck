package view

import (
	_ "embed"

	"github.com/gregoryv/content"
	"github.com/gregoryv/web"
	. "github.com/gregoryv/web"
)

func Present(b *content.Content) *Page {
	body := Body()
	parts := b.Parts()
	if b.Cover() != nil {
		parts = append([]*Element{b.Cover()}, parts...)
	}

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

	p := NewPage(
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
	return p
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
