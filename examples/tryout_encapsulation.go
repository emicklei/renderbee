package main

import (
	"github.com/emicklei/renderbee"
	"html/template"
	"os"
)

type Markup struct {
	Data     interface{}
	Template *template.Template
	Canvas   *renderbee.HtmlCanvas
	Name     string
}

func (c Markup) RenderOn(hc *renderbee.MarkupCanvas) {
	c.Template.Execute(hc, c.Data)
}

func (c Markup) Render() template.HTML {
	return renderbee.HTML(c)
}

type MarkupSequence struct {
	Sequence []Markup
}

var Feed_Template = template.Must(template.New("Feed_Template").Parse(`
<div class="items">
	{{range .Sequence}}
		{{.Render}}
	{{end}}
</div>`))

type FeedItem struct {
	Title       string
	Description string
}

var FeedItem_Template = template.Must(template.New("FeedItem_Template").Parse(`
<div class="item">
	<div class="title">
		{{.Title}}
	</div>
	<div class="desc">
		{{.Description}}
	</div>
</div>`))

var Page_Template = template.Must(template.New("Page_Template").Parse(`
<html>
	<body>
		<h1>Todays news</h1>
		{{.Render "Feed"}}
		<h2>Old news</h2>
		{{.Render "Old"}}
	</body>
</html>
`))

func main() {
	item1 := FeedItem{"Go takes over the world", "The inevitable happened as...."}
	item2 := FeedItem{"Free water for everybody", "A new source of fresh water has been...."}

	c1 := Markup{item1, FeedItem_Template}
	c2 := Markup{item2, FeedItem_Template}

	// NewMarkupSequence
	feed := MarkupSequence{[]Markup{c1, c2}}

	c3 := Markup{feed, Feed_Template}

	// NewMarkupMap
	page := renderbee.NewContainer(Page_Template)
	page.Add("Feed", c3)
	page.Add("Old", c1)

	canvas := renderbee.NewHtmlCanvas(os.Stdout)
	canvas.Render(page)
}
