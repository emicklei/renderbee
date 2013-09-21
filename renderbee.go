package renderbee

import (
	"html/template"
)

// Renderables can render themselves on a HtmlCanvas.
type Renderable interface {
	RenderOn(hc *HtmlCanvas)
}

// Container encapsulates a Template and a set of named Renderable components.
// A Container can be rendered on a HtmlCanvas.
type Container struct {
	components map[string]Renderable
	template   *template.Template
}

// NewContainer creates and returns a Renderable Container with a html Template
func NewContainer(template *template.Template) *Container {
	c := new(Container)
	c.components = map[string]Renderable{}
	c.template = template
	return c
}

// RenderOn executes the template of the Container using the collection
// of named Renderable components.
// Required method from the Renderable interface.
func (c Container) RenderOn(hc *HtmlCanvas) {
	c.template.Execute(hc, componentsRenderer{c.components, hc})
}

// Add a Renderable component with its associated name
// Refer to this component in a template by writing:
// {{.Render "<name>" }}
func (c *Container) Add(name string, r Renderable) {
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
