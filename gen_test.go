package docor

import "testing"

func TestNewFile(t *testing.T) {
	d := NewBlankDoc()
	d.AddTitle("Golang Docor")
	d.AddSection("Introduction")
	d.AddBlockquote("Keep it Simple!")
	d.AddParagraph("Docor is one-page report generator using go, whose semantic is kinda like Tex-thing:")
	d.AddPre(`func (d *Doc) AddPre(content string) {
	d.body.AppendChild(&html.Node{
		Type: html.ElementNode,
		Data: "pre",
		FirstChild: &html.Node{
			Type: html.TextNode,
			Data: content,
		}})
}`)
	d.AddSection("Charts and Images")
	d.AddImgFromUri("https://plotinum.googlecode.com/files/points_commas.png", 3)
	d.AddChartLineXsYs([]float64{1, 2, 3, 4, 5, 6}, []float64{1.24, 2.123, 3.323, 3.9878, 4.143, 5.123}, 8, 4)
	d.SaveTo("index.html")
}
