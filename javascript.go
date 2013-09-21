package renderbee

import "io"

type Javascript struct {
	Href string
}

func (j Javascript) RenderOn(hc *HtmlCanvas) {
	io.WriteString(hc, `<script type="text/javascript" src="`)
	io.WriteString(hc, j.Href)
	io.WriteString(hc, `"></script>`)
}
