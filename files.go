package deck

import (
	"os/exec"
	"strings"

	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/files"
)

func Load(filename string) *Element {
	v := files.MustLoad(filename)
	v = strings.ReplaceAll(v, "\t", "    ")
	v = highlight(v)
	return Div(
		Class("srcfile"),
		Code(numLines(v, 1)),
	)
}

func LoadFunc(filename, funcName string) *Element {
	v := files.MustLoadFunc(filename, funcName)
	v = strings.ReplaceAll(v, "\t", "    ")
	v = highlight(v)
	return Ul(
		Class("srcfile"),
		Code(numLines(v, 1)),
	)
}

func LoadLines(filename string, from, to int) *Element {
	v := files.MustLoadLines(filename, from, to)

	v = strings.ReplaceAll(v, "\t", "    ")
	v = highlight(v)
	return Div(
		Class("srcfile"),
		Code(numLines(v, from)),
	)
}

func numLines(v string, n int) *Element {
	lines := strings.Split(v, "\n")
	ol := Ol()
	if n != 1 {
		ol.With(Attr("start", n))
	}
	for _, l := range lines {
		ol.With(
			Li(l),
		)
	}
	return ol
}

func godoc(pkg, short string) *Element {
	var out []byte
	if short != "" {
		out, _ = exec.Command("go", "doc", short, pkg).Output()
	} else {
		out, _ = exec.Command("go", "doc", pkg).Output()
	}
	v := string(out)
	v = strings.ReplaceAll(v, "\t", "    ")
	v = highlightGoDoc(v)
	return Wrap(
		Pre(v),
		A(Attr("target", "_blank"),
			Href("https://pkg.go.dev/"+pkg),
			"pkg.go.dev/", pkg,
		),
	)
}

func Shell(cmd, filename string) *Element {
	v := files.MustLoad(filename)
	return Pre(Class("shell dark"), cmd, Br(), v)
}
