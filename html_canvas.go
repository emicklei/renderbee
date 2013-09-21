package renderbee

import (
	"io"
)

// HtmlCanvas is a io.Writer for writing Renderable components.
type HtmlCanvas struct {
	io.Writer
}

// Render dispatches to RenderOn on the Renderable component and returns itself.
func (hc *HtmlCanvas) Render(r Renderable) *HtmlCanvas {
	r.RenderOn(hc)
	return hc
}
