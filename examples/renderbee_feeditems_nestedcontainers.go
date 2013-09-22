package main

import (
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
		{{.Render "FeedItem"}}
	{{end}}
</div>`))

func (f Feed) RenderOn(hc *renderbee.HtmlCanvas) {
	Feed_Template.Execute(hc, f)
}

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

func (f FeedItem) RenderOn(hc *renderbee.HtmlCanvas) {
	FeedItem_Template.Execute(hc, f)
}

func (f FeedItem) Render() template.HTML {
	return renderbee.HTML(f)
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

	feedContainer := renderbee.NewContainer(Feed_Template)
	feedContainer.Add("FeedItem", feed)

	page := renderbee.NewContainer(Page_Template)
	page.Add("Feed", feedContainer)

	canvas := renderbee.NewHtmlCanvas(os.Stdout)
	canvas.Render(page)
}
