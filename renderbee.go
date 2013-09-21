package renderbee

import (
	"html/template"
)

type Renderable interface {
	RenderOn(hc *HtmlCanvas)
}

type Container struct {
	Components map[string]Renderable
	Template   *template.Template
}

func NewContainer(template *template.Template) *Container {
	c := new(Container)
	c.Components = map[string]Renderable{}
	c.Template = template
	return c
}

func (c Container) RenderOn(hc *HtmlCanvas) {
	c.Template.Execute(hc, componentsRenderer{c.Components, hc})
}

func (c *Container) Add(name string, r Renderable) {
	c.Components[name] = r
}

type componentsRenderer struct {
	components map[string]Renderable
	canvas     *HtmlCanvas
}

func (c componentsRenderer) Render(name string) string {
	c.components[name].RenderOn(c.canvas)
	return ""
}
