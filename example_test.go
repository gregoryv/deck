package deck_test

import (
	"github.com/gregoryv/deck"
	. "github.com/gregoryv/web"
)

func Example_pack() {
	p := deck.NewPack()
	p.Title = "My funny presentation"
	p.Author = "Gregory Vinčić"
	p.Style(".nowrap", "white-space: nowrap")
	// dark mode
	//p.Style("html, body", "background-color: #2E2E34", "color: #f0f8ff")

	p.NewSlide(
		H2("Animals"),
		deck.Double(
			lorem,
			ipsum(),
		),
	)
	p.NewSlide(
		H3("Horses"),
		deck.Middle(69,
			ipsum().With(Class("nowrap")),
		),
	)
	p.NewSlide(
		H3("Birds"),
		H4("Wings"),
		lorem,
	)
	p.NewSlide(
		H2("Things"),
		lorem,
	)
	p.NewSlide(
		H3("Cars"),
		deck.Center(
			H4("title 4"),
			Img(Src("https://avatars.githubusercontent.com/u/779941?v=4")),
			lorem,
		),
	)

	p.Document().SaveAs("out.html")
	// output:
}

var lorem = P(`Lorem ipsum dolor sit amet, consectetur adipiscing
elit, sed do eiusmod tempor incididunt ut labore et dolore magna
aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco
laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor
in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.`)

func ipsum() *Element {
	return Ul(
		Li("Lorem ipsum dolor sit amet, consectetur adipiscing elit,"),
		Li("sed do eiusmod tempor incididunt ut labore et dolore magna"),
		Li("aliqua. Ut enim ad minim veniam, quis nostrud exercitation"),
		Li("ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis"),
		Li("aute irure dolor in reprehenderit in voluptate velit esse"),
		Li("cillum dolore eu fugiat nulla pariatur. Excepteur sint"),
		Li("occaecat cupidatat non proident, sunt in culpa qui officia"),
		Li("deserunt mollit anim id est laborum."),
	)
}
