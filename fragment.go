package renderbee

import (
	"html/template"
)

// Fragment encapsulates data and a template and can be rendered on a HtmlCanvas.
type Fragment struct {
	data     interface{}
	Template *template.Template
}

// NewFragment creates a Fragment that can be rendered on a HtmlCanvas.
func NewFragment(data interface{}, t *template.Template) *Fragment {
	return &Fragment{data, t}
}

// RenderOn takes the template and data of the fragment to render on a HtmlCanvas.
// part of the Renderable interface
func (c Fragment) RenderOn(hc *HtmlCanvas) {
	c.Template.Execute(hc, c.data)
}

// Return some HTML(string) as a result of the rendering the receiver on a temporary HtmlCanvas.
func (c Fragment) Render() template.HTML {
	return HTML(c)
}
