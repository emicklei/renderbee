package renderbee

import (
	"io"
)

type HtmlCanvas struct {
	io.Writer
}

func (hc HtmlCanvas) Render(r Renderable) {
	r.RenderOn(hc)
}

func (hc HtmlCanvas) Javascript(href string) {
	io.WriteString(hc, `<script type="text/javascript" src="`)
	io.WriteString(hc, href)
	io.WriteString(hc, `"></script>`)
}

func (hc HtmlCanvas) Stylesheet(href string) {
	io.WriteString(hc, `<link type="text/css" rel="stylesheet" href="`)
	io.WriteString(hc, href)
	io.WriteString(hc, `">`)
}
