package renderbee

import (
	"html/template"
)

// FragmentMap encapsulates a Template and a set of named Renderable components.
// A FragmentMap can be rendered on a HtmlCanvas.
type FragmentMap struct {
	components map[string]Renderable
	template   *template.Template
}

// NewFragmentMap creates and returns a Renderable FragmentMap with a html Template.
func NewFragmentMap(template *template.Template) *FragmentMap {
	c := new(FragmentMap)
	c.components = map[string]Renderable{}
	c.template = template
	return c
}

// RenderOn executes the template of the FragmentMap using the collection of named Renderable components.
// Required method from the Renderable interface.
func (c FragmentMap) RenderOn(hc *HtmlCanvas) {
	c.template.Execute(hc, componentsRenderer{c.components, hc})
}

// Put adds a Renderable component with its associated name
// Refer to this component in a template by writing:
// {{.Render "<name>" }}
func (c *FragmentMap) Put(name string, r Renderable) {
	c.components[name] = r
}

// componentsRenderer is a helper class to render components when executing a Template.
type componentsRenderer struct {
	components map[string]Renderable
	canvas     *HtmlCanvas
}

// Render is called because of a template reference: {{.Render "<name>"}}
func (c componentsRenderer) Render(name string) string {
	c.components[name].RenderOn(c.canvas)
	return ""
}
