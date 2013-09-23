package main

import (
	"github.com/emicklei/renderbee"
	"html/template"
	"os"
)

var Feed_Template = template.Must(template.New("Feed_Template").Parse(`
<div class="items">
	{{range .}}
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

	f1 := renderbee.NewFragment(item1, FeedItem_Template)
	f2 := renderbee.NewFragment(item2, FeedItem_Template)

	feed := renderbee.NewFragmentSequence(Feed_Template)
	feed.Add(f1, f2)

	page := renderbee.NewFragmentMap(Page_Template)
	page.Add("Feed", feed) // fragmentsequence
	page.Add("Old", f1)    // fragment

	canvas := renderbee.NewHtmlCanvas(os.Stdout)
	canvas.Render(page)
}
