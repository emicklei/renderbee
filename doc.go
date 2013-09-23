/*
Package renderbee provides a micro-framework for creating HTML components based on the standard html/templates.

Rationale

	Encapsulate a Go template and its data into a Fragment
 	Use a FragmentMap or FragmentSequence to encapsulate a Go template and multiple components
 	Render fragments on a HtmlCanvas (io.Writer) using templates
	In addition to the standard template behavior, you can reference Fragments.

Example

	<html>
	 <head>
	  {{.Render "Head"}
	 </head>
	 <body>
	  <div id="content">
	   {{.Render "Header"}}
	   {{.Render "BodyContent"}}
	   {{.Render "Footer"}}
	  </div>
	 </body>
	</html>

Use this template and fragment components to render a HTML page.

	page := renderBee.NewFragmentMap(Page_Template)
	page.Add("Head", headFragment)
	page.Add("Header", headerFragment)
	page.Add("BodyContent", boydFragmentMap)
	page.Add("Footer", footerFragment)

	canvas := NewHtmlCanvas(resp.ResponseWriter)
	canvas.Render(page)

(c) 2013, http://ernestmicklei.com. Apache V2 License
*/
package renderbee
