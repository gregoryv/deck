package content

import (
	"github.com/gregoryv/content"
	"github.com/gregoryv/content/view"
	. "github.com/gregoryv/web"
)

func ExampleBook() {
	b := content.NewContent()
	title := "My funny presentation"
	b.SetTitle(title)

	author := "Gregory Vinčić"
	b.SetAuthor(author)

	b.SetStyle(myTheme())

	b.NewPart(
		H1("title 1"),

		P(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna
		aliqua. Ut enim ad minim veniam, quis nostrud exercitation
		ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis
		aute irure dolor in reprehenderit in voluptate velit esse
		cillum dolore eu fugiat nulla pariatur. Excepteur sint
		occaecat cupidatat non proident, sunt in culpa qui officia
		deserunt mollit anim id est laborum.`),

		P(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna
		aliqua. Ut enim ad minim veniam, quis nostrud exercitation
		ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis
		aute irure dolor in reprehenderit in voluptate velit esse
		cillum dolore eu fugiat nulla pariatur. Excepteur sint
		occaecat cupidatat non proident, sunt in culpa qui officia
		deserunt mollit anim id est laborum.`),
	)
	b.NewPart(
		H2("Lorem Ipsum"),

		P(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna
		aliqua. Ut enim ad minim veniam, quis nostrud exercitation
		ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis
		aute irure dolor in reprehenderit in voluptate velit esse
		cillum dolore eu fugiat nulla pariatur. Excepteur sint
		occaecat cupidatat non proident, sunt in culpa qui officia
		deserunt mollit anim id est laborum.`),
	)
	b.NewPart(
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
	b.NewPart(
		H4("title 4"),
		"more here",
		P(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna
		aliqua. Ut enim ad minim veniam, quis nostrud exercitation
		ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis
		aute irure dolor in reprehenderit in voluptate velit esse
		cillum dolore eu fugiat nulla pariatur. Excepteur sint
		occaecat cupidatat non proident, sunt in culpa qui officia
		deserunt mollit anim id est laborum.`),
	)

	p := view.NewPresentation(b)
	p.Document().SaveAs("example.html")
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
