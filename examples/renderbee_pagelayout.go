package main

import (
	"github.com/emicklei/renderbee"
	"html/template"
	"os"
)

var PageLayout_Template = template.Must(template.New("PageLayout").Parse(`<html>
<body>
{{.Render "Header"}}
{{.Render "Footer"}}
</body>
</html>
`))

type Header struct {
	Title string
}

func (h Header) RenderOn(hc *renderbee.HtmlCanvas) {
	Header_Template.Execute(hc, h)
}

var Header_Template = template.Must(template.New("Header").Parse(`
<h1>{{.Title}}</h1>
`))

type Footer struct {
	Caption string
}

func (f Footer) RenderOn(hc *renderbee.HtmlCanvas) {
	Footer_Template.Execute(hc, f)
}

var Footer_Template = template.Must(template.New("Footer").Parse(`
<h4>&copy; 2013. All rights reserved. {{.Caption}}</h4>
`))

func main() {
	canvas := renderbee.NewHtmlCanvas(os.Stdout)
	lay := renderbee.NewFragmentMap(PageLayout_Template)
	lay.Put("Header", Header{"Demo renderBee"})
	lay.Put("Footer", Footer{"powered by renderBee"})
	canvas.Render(lay)
}
