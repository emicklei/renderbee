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

func ExampleCompositeFragment() {
	citiesTemplate := template.Must(template.New("Page").Parse(`
<div class="cities">
{{range .}}{{.Render}}
{{end}}</div>`))
	cityTemplate := template.Must(template.New("City").Parse(
		`<h1>Welcome to {{.}}</h1>`))
	fragment := NewCompositeFragment(citiesTemplate)
	fragment.Add(NewFragment("New York", cityTemplate))
	fragment.Add(NewFragment("Amsterdam", cityTemplate))
	canvas := NewHtmlCanvas(os.Stdout)
	canvas.Render(fragment)
	// Output:
	//<div class="cities">
	//<h1>Welcome to New York</h1>
	//<h1>Welcome to Amsterdam</h1>
	//</div>
}

func ExampleFragmentMap() {
	pageTemplate := template.Must(template.New("Page").Parse(`
<html>
<body>
{{.Render "Header"}}
{{.Render "Footer"}}
</body>
</html>`))
	headerTemplate := template.Must(template.New("Header").Parse(
		`<h1>Hello {{.}}</h1>`))
	footerTemplate := template.Must(template.New("Footer").Parse(
		`<h1>&copy; 2013. Created by {{.}}</h1>`))
	page := NewFragmentMap(pageTemplate)
	page.Put("Header", NewFragment("John", headerTemplate))
	page.Put("Footer", NewFragment("Jane", footerTemplate))
	canvas := NewHtmlCanvas(os.Stdout)
	canvas.Render(page)
	// Output:
	//<html>
	//<body>
	//<h1>Hello John</h1>
	//<h1>&copy; 2013. Created by Jane</h1>
	//</body>
	//</html>
}
