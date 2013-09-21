package renderbee

import "io"

type Stylesheet struct {
	Href string
}

func (s Stylesheet) RenderOn(hc *HtmlCanvas) {
	io.WriteString(hc, `<link type="text/css" rel="stylesheet" href="`)
	io.WriteString(hc, s.Href)
	io.WriteString(hc, `"/>`)
}
