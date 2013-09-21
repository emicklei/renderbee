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

// Return a new HtmlCanvas for writing Html on an io.Writer
func NewHtmlCanvas(writer io.Writer) *HtmlCanvas {
	hc := new(HtmlCanvas)
	hc.Writer = writer
	return hc
}
