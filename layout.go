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

func layoutView() *CSS {
	css := NewCSS()
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
