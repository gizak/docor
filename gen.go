package docor

import "golang.org/x/net/html"
import "bytes"
import "github.com/GeertJohan/go.rice"
import "code.google.com/p/plotinum/plot"
import "code.google.com/p/plotinum/plotter"
import "code.google.com/p/plotinum/plotutil"
import "strconv"
import "os"
import "encoding/base64"
import "io/ioutil"

type Doc struct {
	buf  *bytes.Buffer
	root *html.Node
	body *html.Node
	head *html.Node
}

func newRawDoc() *Doc {
	d := &Doc{}
	d.buf = new(bytes.Buffer)
	d.buf.WriteString("<!doctype html><html><head></head><body><div class='container'></div></body></html>")
	root, err := html.Parse(d.buf)
	if err != nil {
		panic(err)
	}
	d.root = root
	d.body = d.root.LastChild.LastChild.FirstChild
	d.head = d.root.LastChild.FirstChild
	return d
}

func NewBlankDoc() *Doc {
	d := newRawDoc()
	css := rice.MustFindBox("resources/css").MustString("materialize.css")
	cssCodeHighlight := rice.MustFindBox("resources/css").MustString("prism.css")

	d.head.AppendChild(
		&html.Node{
			Type: html.ElementNode,
			Data: "style",
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: css,
			}})
	d.head.AppendChild(
		&html.Node{
			Type: html.ElementNode,
			Data: "style",
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: cssCodeHighlight,
			}})

	jsCodeHighlight := rice.MustFindBox("resources/js").MustString("prism.js")
	d.head.AppendChild(textNode("script", jsCodeHighlight, genAttrs("type", "text/javascript")))

	return d
}

func genAttrs(as ...string) []html.Attribute {
	l := len(as) / 2
	attrs := make([]html.Attribute, l)
	for i := range attrs {
		attrs[i] = html.Attribute{"", as[2*i], as[2*i+1]}
	}
	return attrs
}

func textNode(tag, content string, attrs []html.Attribute) *html.Node {
	return &html.Node{
		Type: html.ElementNode,
		Data: tag,
		Attr: attrs,
		FirstChild: &html.Node{
			Type: html.TextNode,
			Data: content}}
}

func (d *Doc) AddTitle(title string) {
	d.body.AppendChild(textNode("h2", title, genAttrs("class", "center")))
}

func (d *Doc) AddSection(title string) {
	d.body.AppendChild(textNode("h3", title, nil))
}

func (d *Doc) AddSubsection(title string) {
	d.body.AppendChild(textNode("h4", title, nil))
}

func (d *Doc) AddPre(content string) {
	d.body.AppendChild(textNode("pre", content, nil))
}

func (d *Doc) AddParagraph(par string) {
	d.body.AppendChild(textNode("p", par, nil))
}

func (d *Doc) AddBlockquote(quote string) {
	d.body.AppendChild(textNode("blockquote", quote, nil))
}

func rowCardNode(width int) *html.Node {
	div := &html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: genAttrs("class", "card")}
	row := &html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: genAttrs("class", "row"),
		FirstChild: &html.Node{
			Type:       html.ElementNode,
			Attr:       genAttrs("class", "col s12 m"+strconv.Itoa(width)),
			Data:       "div",
			FirstChild: div,
		}}
	return row
}

func (d *Doc) AddImgFromURI(url string, width int) {
	row := rowCardNode(width)
	row.FirstChild.FirstChild.AppendChild(&html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: genAttrs("class", "card-image"),
		FirstChild: &html.Node{
			Type: html.ElementNode,
			Data: "img",
			Attr: genAttrs("src", url),
		}})

	d.body.AppendChild(row)
}

func (d *Doc) AddCode(code, lang string) {
	pre := &html.Node{
		Type: html.ElementNode,
		Data: "pre",
		FirstChild: &html.Node{
			Type: html.ElementNode,
			Data: "code",
			Attr: genAttrs("class", "language-"+lang),
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: code,
			}}}
	d.body.AppendChild(pre)
}

func (d *Doc) AddDetailImgFromURI(url string, width int, title, titleColor, description string) {
	row := rowCardNode(width)
	row.FirstChild.FirstChild.AppendChild(&html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: genAttrs("class", "card-image")})

	row.FirstChild.FirstChild.FirstChild.AppendChild(&html.Node{
		Type: html.ElementNode,
		Data: "img",
		Attr: genAttrs("src", url)})

	row.FirstChild.FirstChild.FirstChild.AppendChild(textNode("span", title, genAttrs("class", "card-title "+titleColor+"-text")))

	row.FirstChild.FirstChild.AppendChild(&html.Node{
		Type:       html.ElementNode,
		Data:       "div",
		Attr:       genAttrs("class", "card-content"),
		FirstChild: textNode("p", description, nil)})
	d.body.AppendChild(row)
}

func (d *Doc) AddChartLineXsYs(xs, ys []float64, w, h int) {
	p, _ := plot.New()
	//p.X.Label.Text = "X"
	//p.Y.Label.Text = "Y"
	xys := make(plotter.XYs, len(xs))
	for i := range xys {
		xys[i].X = xs[i]
		xys[i].Y = ys[i]
	}
	plotutil.AddLinePoints(p, "", xys)
	tmp := os.TempDir()
	f := tmp + "/" + "svgcache.svg"
	if err := p.Save(float64(w), float64(h), f); err != nil {
		panic(err)
	} else {
		b, _ := ioutil.ReadFile(f)
		b64 := base64.StdEncoding.EncodeToString(b)
		svgb64 := "data:image/svg+xml;charset=utf-8;base64," + b64
		row := rowCardNode(w)
		row.FirstChild.FirstChild.AppendChild(&html.Node{
			Type: html.ElementNode,
			Data: "div",
			Attr: genAttrs("class", "card-image", "style", "padding-top:10px;padding-left:10px"),
			FirstChild: &html.Node{
				Type: html.ElementNode,
				Data: "img",
				Attr: genAttrs("src", svgb64),
			}})
		d.body.AppendChild(row)
	}
}

func (d *Doc) AddTable(thead []string, tbody [][]string) {
	tbl := &html.Node{
		Type: html.ElementNode,
		Data: "table",
		Attr: genAttrs("class", "hoverable"),
	}

	th := &html.Node{
		Type: html.ElementNode,
		Data: "thead",
		FirstChild: &html.Node{
			Type: html.ElementNode,
			Data: "tr",
		}}
	for _, v := range thead {
		th.FirstChild.AppendChild(textNode("th", v, nil))
	}

	tb := &html.Node{
		Type: html.ElementNode,
		Data: "tbody",
	}

	for i := range tbody {
		tr := &html.Node{
			Type: html.ElementNode,
			Data: "tr",
		}
		for _, v := range tbody[i] {
			tr.AppendChild(textNode("td", v, nil))
		}
		tb.AppendChild(tr)
	}

	tbl.AppendChild(th)
	tbl.AppendChild(tb)
	d.body.AppendChild(tbl)
}

func (d *Doc) SaveTo(fname string) error {

	html.Render(d.buf, d.root)
	return ioutil.WriteFile(fname, d.buf.Bytes(), 0777)
}
