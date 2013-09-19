package renderbee

import (
	"html/template"
)

type Renderable interface {
	RenderOn(h HtmlCanvas)
}

type Decorate struct {
	Component Renderable
}

func (d Decorate) Execute(t *template.Template, data interface{}, hc HtmlCanvas) {
	t.Execute(hc, componentRenderer{d.Component, hc})
}

type componentRenderer struct {
	component Renderable
	canvas    HtmlCanvas
}

func (c componentRenderer) Render() string {
	c.component.RenderOn(c.canvas)
	return ""
}
