package content

import (
	"github.com/gregoryv/content"
	"github.com/gregoryv/content/view"
	. "github.com/gregoryv/web"
)

func ExampleBook() {
	c := content.NewContent()
	c.SetTitle("My funny presentation")
	c.SetAuthor("Gregory Vinčić")
	c.SetStyle(myTheme())

	p := view.NewPresentation(c)
	p.NewSlide(
		H2("Here we go"),

		P(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna
		aliqua. Ut enim ad minim veniam, quis nostrud exercitation
		ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis
		aute irure dolor in reprehenderit in voluptate velit esse
		cillum dolore eu fugiat nulla pariatur. Excepteur sint
		occaecat cupidatat non proident, sunt in culpa qui officia
		deserunt mollit anim id est laborum.`),

		Ul(
			Li("Lorem ipsum dolor sit amet, consectetur adipiscing elit,"),
			Li("sed do eiusmod tempor incididunt ut labore et dolore magna"),
			Li("aliqua. Ut enim ad minim veniam, quis nostrud exercitation"),
			Li("ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis"),
			Li("aute irure dolor in reprehenderit in voluptate velit esse"),
			Li("cillum dolore eu fugiat nulla pariatur. Excepteur sint"),
			Li("occaecat cupidatat non proident, sunt in culpa qui officia"),
			Li("deserunt mollit anim id est laborum."),
		),
	)
	p.NewSlide(
		H2("Lorem Ipsum"),

		Ul(Class("nowrap"),
			Li("Lorem ipsum dolor sit amet, consectetur adipiscing elit,"),
			Li("sed do eiusmod tempor incididunt ut labore et dolore magna"),
			Li("aliqua. Ut enim ad minim veniam, quis nostrud exercitation"),
			Li("ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis"),
			Li("aute irure dolor in reprehenderit in voluptate velit esse"),
			Li("cillum dolore eu fugiat nulla pariatur. Excepteur sint"),
			Li("occaecat cupidatat non proident, sunt in culpa qui officia"),
			Li("deserunt mollit anim id est laborum."),
		),
	)

	p.NewSlide(
		H2("Third"),
		H3("title 3"),
		P(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna
		aliqua. Ut enim ad minim veniam, quis nostrud exercitation
		ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis
		aute irure dolor in reprehenderit in voluptate velit esse
		cillum dolore eu fugiat nulla pariatur. Excepteur sint
		occaecat cupidatat non proident, sunt in culpa qui officia
		deserunt mollit anim id est laborum.`),
	)
	p.NewSlide(
		H2("Fourth"),
		Span(
			H4("title 4"),
			"more here",
		),
		P(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna
		aliqua. Ut enim ad minim veniam, quis nostrud exercitation
		ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis
		aute irure dolor in reprehenderit in voluptate velit esse
		cillum dolore eu fugiat nulla pariatur. Excepteur sint
		occaecat cupidatat non proident, sunt in culpa qui officia
		deserunt mollit anim id est laborum.`),
	)

	p.Document().SaveAs("presentation.html")
	// output:
}

func myTheme() *CSS {
	css := NewCSS()
	css.Style("html, body",
		"background-color: #999999",
	)
	css.Style(".page .content",
		"background-color: #fff",
	)
	css.Style(".page .footer",
		"background-color: #fff",
	)
	css.Style("h1",
		"text-align: center",
	)
	return css
}
