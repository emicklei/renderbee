package main

import (
	"github.com/emicklei/renderbee"
	"html/template"
	"io"
	"os"
)

type PageLayout struct {
	renderbee.Decorate
}

func (p PageLayout) RenderOn(w io.Writer) {
	p.Decorate.Execute(PageLayout_Template, p, w)
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

func (h Header) RenderOn(w io.Writer) {
	Header_Template.Execute(w, h)
}

var Header_Template = template.Must(template.New("Header").Parse(`
<h1>{{.Title}}</h1>
`))

func main() {
	PageLayout{renderbee.Decorate{Header{"renderBee"}}}.RenderOn(os.Stdout)
}
