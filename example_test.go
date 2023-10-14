package deck_test

import (
	"github.com/gregoryv/deck"
	. "github.com/gregoryv/web"
)

func Example_deck() {
	p := deck.Deck{
		Title:     "My funny presentation",
		Author:    "Gregory Vinčić",
		AutoCover: true,
		AutoTOC:   true,
	}
	// dark mode
	//p.Style("html, body, a", "background-color: #2E2E34", "color: #f0f8ff")
	p.Style(".header", "background-color: #e2e2e2")
	// center toc if short titles
	//p.Style(".toc", "margin-left: 10vw", "width: 20vw")

	p.NewCard(
		// each card must start with H2 or H3
		H2("Animals"),
		deck.Double(
			lorem,
			ipsum,
		),
	)
	p.NewCard(
		// H3 means sub card of last H"
		H3("Horses"),
		deck.Middle(50,
			ipsum,
		),
	)
	p.NewCard(
		H3("Birds"),
		H4("Wings"),
		lorem,
	)
	p.NewCard(
		H2("Things"),
		lorem,
	)
	p.NewCard(
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

var ipsum = Ul(
	Li("Lorem ipsum dolor sit amet, consectetur adipiscing elit,"),
	Li("sed do eiusmod tempor incididunt ut labore et dolore magna"),
	Li("aliqua. Ut enim ad minim veniam, quis nostrud exercitation"),
	Li("ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis"),
	Li("aute irure dolor in reprehenderit in voluptate velit esse"),
	Li("cillum dolore eu fugiat nulla pariatur. Excepteur sint"),
	Li("occaecat cupidatat non proident, sunt in culpa qui officia"),
	Li("deserunt mollit anim id est laborum."),
)
