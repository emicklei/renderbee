package renderbee

import "os"

func ExampleJavascript() {
	canvas := HtmlCanvas{os.Stdout}
	canvas.Render(Javascript{"http://mysite.com/static/main.js"})
	// Output:
	// <script type="text/javascript" src="http://mysite.com/static/main.js"></script>
}

func ExampleStylesheet() {
	canvas := HtmlCanvas{os.Stdout}
	canvas.Render(Stylesheet{"http://mysite.com/static/main.css"})
	// Output:
	// <link type="text/css" rel="stylesheet" href="http://mysite.com/static/main.css"/>
}
