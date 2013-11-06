package renderbee

import (
	"html/template"
)

// CompositeFragment encapsulates a slice of Fragment and a template and can be rendered on a HtmlCanvas.
type CompositeFragment struct {
	Fragments []*Fragment
	Template  *template.Template
}

// NewCompositeFragment creates an empty CompositeFragment with a template.
func NewCompositeFragment(t *template.Template) *CompositeFragment {
	seq := new(CompositeFragment)
	seq.Template = t
	seq.Fragments = []*Fragment{}
	return seq
}

// Add a Fragement to the slice of Fragments.
func (c *CompositeFragment) Add(frags ...*Fragment) {
	c.Fragments = append(c.Fragments, frags...)
}

// RenderOn takes the template and collection of Fragments to render on a HtmlCanvas.
// part of the Renderable interface
func (c CompositeFragment) RenderOn(hc *HtmlCanvas) {
	c.Template.Execute(hc, c.Fragments)
}

// Return some HTML(string) as a result of the rendering the receiver on a temporary HtmlCanvas.
func (c CompositeFragment) Render() template.HTML {
	return HTML(c)
}
