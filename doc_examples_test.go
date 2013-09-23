package renderbee

import (
	"html/template"
	"os"
)

func ExampleJavascript() {
	canvas := NewHtmlCanvas(os.Stdout)
	canvas.Render(Javascript{"http://mysite.com/static/main.js"})
	// Output:
	// <script type="text/javascript" src="http://mysite.com/static/main.js"></script>
}

func ExampleStylesheet() {
	canvas := NewHtmlCanvas(os.Stdout)
	canvas.Render(Stylesheet{"http://mysite.com/static/main.css"})
	// Output:
	// <link type="text/css" rel="stylesheet" href="http://mysite.com/static/main.css"/>
}

func ExampleFragment() {
	person := struct {
		Name  string
		Email string
	}{
		"John",
		"john.doe@here.com",
	}
	personTemplate := template.Must(template.New("Person_Template").Parse(`
<div class="person">
	<a href="{{.Email}}">{{.Name}}</a>
</div>`))
	fragment := NewFragment(person, personTemplate)
	canvas := NewHtmlCanvas(os.Stdout)
	canvas.Render(fragment)
	// Output:
	// <div class="person">
	//	<a href="john.doe@here.com">John</a>
	// </div>
}
