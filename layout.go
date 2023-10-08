package deck

import (
	"fmt"

	. "github.com/gregoryv/web"
)

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

func Middle(width int, element ...any) *Element {
	left := (100 - width) / 2
	div := Div(
		Attr(
			"style",
			fmt.Sprintf(
				"position: absolute; left: %s; width: %s", vw(left), vw(width),
			),
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

func layoutView() *CSS {
	css := NewCSS()
	css.Style(".outline",
		"border: 1px solid black",
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
