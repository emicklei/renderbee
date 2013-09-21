package main

import (
	"github.com/emicklei/renderbee"
	"html/template"
	"os"
)

var PageLayout_Template = template.Must(template.New("PageLayout").Parse(`
<html>
<body>
{{.Render "Header"}}
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

func main() {
	hc := renderbee.HtmlCanvas{os.Stdout}
	lay := renderbee.NewContainer(PageLayout_Template)
	lay.Add("Header", &Header{"renderBee"})
	hc.Render(lay)
}
