package deck

import (
	"regexp"

	. "github.com/gregoryv/web"
)

// highlight go source code
func highlight(v string) string {
	v = keywords.ReplaceAllString(v, `$1<span class="keyword">$2</span>$3`)
	v = types.ReplaceAllString(v, `$1<span class="type">$2</span>$3`)
	v = comments.ReplaceAllString(v, `<span class="comment">$1</span>`)
	return v
}

// highlightGoDoc output
func highlightGoDoc(v string) string {
	v = docKeywords.ReplaceAllString(v, `$1<span class="keyword">$2</span>$3`)
	v = types.ReplaceAllString(v, `$1<span class="type">$2</span>$3`)
	v = comments.ReplaceAllString(v, `<span class="comment">$1</span>`)
	return v
}

var types = regexp.MustCompile(`(\W)(\w+\.\w+)(\)|\n)`)
var keywords = regexp.MustCompile(`(\W?)(^package|const|select|defer|import|for|func|range|return|go|var|switch|if|case|label|type|struct|interface|default)(\W)`)
var docKeywords = regexp.MustCompile(`(\W?)(^package|func|type|struct|interface)(\W)`)
var comments = regexp.MustCompile(`(//[^\n]*)`)

func highlightColors() *CSS {
	css := NewCSS()
	css.Style(".keyword", "color: darkviolet")
	css.Style(".type", "color: dodgerblue")
	css.Style(".comment, .comment>span", "color: darkgreen")
	return css
}

func srcfile(fontSize int) *CSS {
	css := NewCSS()
	css.Style(".srcfile",
		"margin-top: 1.6em",
		"margin-bottom: 1.6em",
		"background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABQAAAAlCAMAAAB1cTk3AAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAjVBMVEX9/fz4+Pb19fP8/Pz39/Xp6ejr6+rs7Ov29vTt7evo6Oj6+vr7+/vv7+/u7u3v7+37+/ru7uzw8PD////29vbx8fD6+vjq6unz8/L5+ff39/f9/f3y8vHo6Ofx8fHw8O/z8/P6+vn4+Pj09PL+/v729vXp6efy8vDv7+709PTz8/H7+/n8/Pv5+fju7u5JBELWAAAAAWJLR0QB/wIt3gAAAAlwSFlzAAALEwAACxMBAJqcGAAAAAd0SU1FB+cFHw8tIWAQRK8AAAC5SURBVCjPpZDbEoIgEIaXQLPUNNFOWJrRyQ7v/3gJAs2EFzTuBex+8+/uDwDIiglgGyItJJ4/9YM+DxSczcMowvFCFgmk4lpmVEZeiGolYbCmKjZiQgp4i5C305CVPezOPTVxMLD6wrorj79K0U6kpaQxM0/GPIm5gmdilKise9ZI4UW96HoLGcN3ybpF+hPaomxV+oCn/Ukv4DZMgdqQOkMMb0clG4L5uO3/wAGf3Ll9UJk5Lxpr/gOpxRF8cA+lxgAAACV0RVh0ZGF0ZTpjcmVhdGUAMjAyMy0wNS0zMVQxMzo0NTozMyswMjowMJXHbJYAAAAldEVYdGRhdGU6bW9kaWZ5ADIwMjMtMDUtMzFUMTM6NDU6MzMrMDI6MDDkmtQqAAAAAElFTkSuQmCC)",
		"background-repeat: repeat-y",
		"padding-left: 36px",
		"background-color: #fafafa",
		"tab-size: 4",
		"-moz-tab-size: 4",
		"min-width: 35vw",
		"font-size: "+vh(fontSize),
	)

	css.Style(".srcfile code",
		"padding: .6em 0 2vh 0",
		"background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABQAAAAlCAMAAAB1cTk3AAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAjVBMVEX9/fz4+Pb19fP8/Pz39/Xp6ejr6+rs7Ov29vTt7evo6Oj6+vr7+/vv7+/u7u3v7+37+/ru7uzw8PD////29vbx8fD6+vjq6unz8/L5+ff39/f9/f3y8vHo6Ofx8fHw8O/z8/P6+vn4+Pj09PL+/v729vXp6efy8vDv7+709PTz8/H7+/n8/Pv5+fju7u5JBELWAAAAAWJLR0QB/wIt3gAAAAlwSFlzAAALEwAACxMBAJqcGAAAAAd0SU1FB+cFHw8tIWAQRK8AAAC5SURBVCjPpZDbEoIgEIaXQLPUNNFOWJrRyQ7v/3gJAs2EFzTuBex+8+/uDwDIiglgGyItJJ4/9YM+DxSczcMowvFCFgmk4lpmVEZeiGolYbCmKjZiQgp4i5C305CVPezOPTVxMLD6wrorj79K0U6kpaQxM0/GPIm5gmdilKise9ZI4UW96HoLGcN3ybpF+hPaomxV+oCn/Ukv4DZMgdqQOkMMb0clG4L5uO3/wAGf3Ll9UJk5Lxpr/gOpxRF8cA+lxgAAACV0RVh0ZGF0ZTpjcmVhdGUAMjAyMy0wNS0zMVQxMzo0NTozMyswMjowMJXHbJYAAAAldEVYdGRhdGU6bW9kaWZ5ADIwMjMtMDUtMzFUMTM6NDU6MzMrMDI6MDDkmtQqAAAAAElFTkSuQmCC)",
		"background-repeat: repeat-y",
		"background-position: right top",
		"display: block",
		"text-align: left",
		"font-family: Inconsolata, monospace",
		"overflow: hidden",
	)
	css.Style(".srcfile ol",
		"margin-left: 1vw",
	)
	css.Style(".srcfile ol li",
		"white-space: pre",
	)
	css.Style(".srcfile ol li::marker",
		"color: RGB(0,0,0,0.2)",
	)
	css.Style(".srcfile ol li:hover",
		//		"background-color: #b4eeb4",
		"background-color: RGB(180,238,180,0.3)",
	)

	css.Style(".srcfile code i",
		"font-style: normal",
		"color: #a2a2a2",
	)
	return css
}
