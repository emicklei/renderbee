package main

import (
	"github.com/emicklei/renderbee"
	"html/template"
	"os"
)

type PageLayout struct {
	renderbee.Decorate
}

func (p PageLayout) RenderOn(hc renderbee.HtmlCanvas) {
	p.Decorate.Execute(PageLayout_Template, p, hc)
}

var PageLayout_Template = template.Must(template.New("PageLayout").Parse(`
<html>
<body>
{{.Render}}
</body>
</html>
`))

type Header struct {
	Title string
}

func (h Header) RenderOn(hc renderbee.HtmlCanvas) {
	Header_Template.Execute(hc, h)
}

var Header_Template = template.Must(template.New("Header").Parse(`
<h1>{{.Title}}</h1>
`))

func main() {
	hc := renderbee.HtmlCanvas{os.Stdout}
	layout := PageLayout{renderbee.Decorate{Header{"renderBee"}}}
	hc.Render(layout)
}
