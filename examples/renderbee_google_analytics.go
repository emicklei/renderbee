package main

import (
	"github.com/emicklei/renderbee"
	"html/template"
	"os"
)

// GoogleAnalytics is a simple compent that encapsulates a HTML script node for
// communicating with google-analytics.com.
// It adds its own implementation of the Renderable interface such that it can be used
// in
type GoogleAnalytics struct {
	Code string
}

// RenderOn writes the result of executing its template with the receiver as its data
// implements renderbee.Renderable so it can be used in a Container.
//
func (g GoogleAnalytics) RenderOn(hc *renderbee.HtmlCanvas) {
	GoogleAnalytics_Template.Execute(hc, g)
}

var GoogleAnalytics_Template = template.Must(template.New("GoogleAnalytics").Parse(`
<script>
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
ga('create', '{{.Code}}');
ga('send', 'pageview');
</script>
`))

func main() {
	canvas := renderbee.NewHtmlCanvas(os.Stdout)
	canvas.Render(GoogleAnalytics{"UA999"})
}
