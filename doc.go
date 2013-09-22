/*
Package renderbee provides a micro-framework for creating HTML components based on the standard html/templates.

Rationale

	Encapsulate a Go template and its data into a component
 	Use a Container to encapsulate a Go template and multiple components
 	Render components and containers on a HtmlCanvas (io.Writer)

Example

	type FeedItem struct {
		Title string
		Description
	}

	var FeedItem_Template = template.Must(template.New("FeedItem_Template").Parse(`
	<div class="item">
		<div class="title">
			{{.Title}}
		</div>
		<div class="desc">
			{{.Description}}
		</div>
	</div>
	`))

	func (f FeedItem) RenderOn(hc *renderbee.HtmlCanvas) {
		FeedItem_Template.Execute(hc, f)
	}

FeedItem is a component that encapsulats data (title,description) and a static template (FeedItem_Template).
In isolation, this component can be rendered (written in HTML on a io.Writer) using:

	canvas := renderbee.NewHtmlCanvas(os.Stdout)
	canvas.Render(FeedItem{"Go takes over the world", "The inevitable happened as...."})

(c) 2013, http://ernestmicklei.com. Apache V2 License
*/
package renderbee
