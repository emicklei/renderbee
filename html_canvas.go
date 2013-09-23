package renderbee

import (
	"bytes"
	"html/template"
	"io"
)

// Renderables can render themselves on a HtmlCanvas.
type Renderable interface {
	RenderOn(hc *HtmlCanvas)
}

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

// HTML returns a template.HTML from rendering a component r.
func HTML(r Renderable) template.HTML {
	buffer := new(bytes.Buffer)
	canvas := NewHtmlCanvas(buffer)
	r.RenderOn(canvas)
	return template.HTML(buffer.String())
}
