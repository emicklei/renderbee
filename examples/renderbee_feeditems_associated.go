package main

import (
	//	"fmt"
	"github.com/emicklei/renderbee"
	"html/template"
	"os"
)

type Feed struct {
	Items []FeedItem
}

var Feed_Template = template.Must(template.New("Feed_Template").Parse(`
<div class="items">
	{{range .Items}}
		{{template "FeedItem_Template" .}}
	{{end}}
</div>

{{define "FeedItem_Template"}}
<div class="item">
	<div class="title">
		{{.Title}}
	</div>
	<div class="desc">
		{{.Description}}
	</div>
</div>
{{end}}`))

func (f Feed) RenderOn(hc *renderbee.HtmlCanvas) {
	Feed_Template.Execute(hc, f)
}

type FeedItem struct {
	Title       string
	Description string
}

var Page_Template = template.Must(template.New("Page_Template").Parse(`
<html>
	<body>
		<h1>Todays news</h1>
		{{.Render "Feed"}}
	</body>
</html>
`))

func main() {
	item1 := FeedItem{"Go takes over the world", "The inevitable happened as...."}
	item2 := FeedItem{"Free water for everybody", "A new source of fresh water has been...."}
	feed := Feed{[]FeedItem{item1, item2}}

	page := renderbee.NewFragmentMap(Page_Template)
	page.Put("Feed", feed)

	canvas := renderbee.NewHtmlCanvas(os.Stdout)
	canvas.Render(page)
}
