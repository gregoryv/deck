package view_test

import (
	"github.com/gregoryv/view"
	. "github.com/gregoryv/web"
)

func ExamplePresentation() {
	p := view.NewPresentation(3)
	p.Title = "My funny presentation"
	p.Author = "Gregory Vinčić"
	p.Style(".nowrap",
		"white-space: nowrap",
	)
	// dark mode
	//p.Style("html, body", "background-color: #2E2E34", "color: #f0f8ff")

	p.NewSlide(
		H2("Here we go"),

		view.Double(
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

		view.Middle(30,
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

		H3("sub title title 3"),
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

		view.Center(
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

	p.Document().SaveAs("example/presentation.html")
	// output:
}
