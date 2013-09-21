package renderbee

import "io"

// Javascript is a simple Renderable components that encapsulates
// the HTML to reference a Javascript file by its hyperlink reference.
type Javascript struct {
	Href string
}

// RenderOn writes the HTML snippet for an external Javascript reference on a HtmlCanvas.
func (j Javascript) RenderOn(hc *HtmlCanvas) {
	io.WriteString(hc, `<script type="text/javascript" src="`)
	io.WriteString(hc, j.Href)
	io.WriteString(hc, `"></script>`)
}

// Stylesheet is a simple Renderable components that encapsulates
// the HTML to reference a Stylesheet file by its hyperlink reference.
type Stylesheet struct {
	Href string
}

// RenderOn writes the HTML snippet for an external Stylesheet reference on a HtmlCanvas.
func (s Stylesheet) RenderOn(hc *HtmlCanvas) {
	io.WriteString(hc, `<link type="text/css" rel="stylesheet" href="`)
	io.WriteString(hc, s.Href)
	io.WriteString(hc, `"/>`)
}
