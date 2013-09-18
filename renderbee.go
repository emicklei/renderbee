package renderbee

import (
	"html/template"
	"io"
)

type Renderable interface {
	RenderOn(w io.Writer)
}

type Decorate struct {
	Component Renderable
}

func (d Decorate) Execute(t *template.Template, data interface{}, w io.Writer) {
	t.Execute(w, ComponentRenderer{d.Component, w})
}

type ComponentRenderer struct {
	Component Renderable
	Writer    io.Writer
}

func (c ComponentRenderer) Render() string {
	c.Component.RenderOn(c.Writer)
	return ""
}
