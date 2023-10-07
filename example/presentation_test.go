package content

import (
	"fmt"

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

		Double(
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
		),
	)
	p.NewSlide(
		H2("Middle No wrap"),

		Middle(30,
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

		Center(
			H4("title 4"),
			Img(Src("https://avatars.githubusercontent.com/u/779941?v=4")),

			P(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna
		aliqua. Ut enim ad minim veniam, quis nostrud exercitation
		ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis
		aute irure dolor in reprehenderit in voluptate velit esse
		cillum dolore eu fugiat nulla pariatur. Excepteur sint
		occaecat cupidatat non proident, sunt in culpa qui officia
		deserunt mollit anim id est laborum.`),
		),
	)

	p.Document().SaveAs("presentation.html")
	// output:
}

func Double(e1, e2 any) *Element {
	div := Div(Class("double"))
	div.With(
		Div(
			Class("column left"),
			e1,
		),
	)
	div.With(
		Div(
			Class("column right"),
			e2,
		),
	)
	return div
}

func Middle(leftVH int, element ...any) *Element {
	div := Div(
		Attr(
			"style",
			fmt.Sprintf("padding-left: %vvh; width: %vvh", leftVH, 100-2*leftVH),
		),
	)
	div.With(element...)
	return div
}

func Center(element ...any) *Element {
	div := Div(Class("center"))
	div.With(element...)
	return div
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
	css.Style(".nowrap",
		"white-space: nowrap",
	)

	// center
	css.Style(".center",
		"text-align: center",
	)
	// double
	css.Style(".double")
	css.Style(".column.left",
		"position: absolute",
		"left: 4vw",
		"width: 40vw",
	)
	css.Style(".column.right",
		"position: absolute",
		"left: 50vw",
		"width: 40vw",
	)
	return css
}
