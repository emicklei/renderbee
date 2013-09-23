package renderbee

import (
	"html/template"
)

// Fragment encapsulates data and a template and can be rendered on a HtmlCanvas.
type Fragment struct {
	data     interface{}
	template *template.Template
}

// NewFragment creates a Fragment that can be rendered on a HtmlCanvas.
func NewFragment(data interface{}, t *template.Template) *Fragment {
	return &Fragment{data, t}
}

// RenderOn takes the template and data of the fragment to render on a HtmlCanvas.
// part of the Renderable interface
func (c Fragment) RenderOn(hc *HtmlCanvas) {
	c.template.Execute(hc, c.data)
}

// Return some HTML(string) as a result of the rendering the receiver on a temporary HtmlCanvas.
func (c Fragment) Render() template.HTML {
	return HTML(c)
}

// FragmentSequence encapsulates a slice of Fragment and a template and can be rendered on a HtmlCanvas.
type FragmentSequence struct {
	Fragments []*Fragment
	Template  *template.Template
}

// NewFragmentSequence creates an empty FragmentSequence with a template.
func NewFragmentSequence(t *template.Template) *FragmentSequence {
	seq := new(FragmentSequence)
	seq.Template = t
	seq.Fragments = []*Fragment{}
	return seq
}

// Add a Fragement to the slice of Fragments.
func (f *FragmentSequence) Add(frags ...*Fragment) {
	f.Fragments = append(f.Fragments, frags...)
}

// RenderOn takes the template and collection of Fragments to render on a HtmlCanvas.
// part of the Renderable interface
func (f FragmentSequence) RenderOn(hc *HtmlCanvas) {
	f.Template.Execute(hc, f.Fragments)
}

// Return some HTML(string) as a result of the rendering the receiver on a temporary HtmlCanvas.
func (f FragmentSequence) Render() template.HTML {
	return HTML(f)
}
