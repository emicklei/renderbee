package renderbee

import (
	"io"
)

type HtmlCanvas struct {
	io.Writer
}

func (hc *HtmlCanvas) Render(r Renderable) *HtmlCanvas {
	r.RenderOn(hc)
	return hc
}
