package docor

import "testing"

func TestNewFile(t *testing.T) {
	d := NewBlankDoc()
	d.AddTitle("Golang Docor")
	d.AddSection("Introduction")
	d.AddBlockquote("Keep it Simple!")
	d.AddParagraph("Docor is one-page report generator using go, whose semantic is kinda like Tex-thing:")
	d.AddCode(`func (d *Doc) AddPre(content string) {
	d.body.AppendChild(&html.Node{
		Type: html.ElementNode,
		Data: "pre",
		FirstChild: &html.Node{
			Type: html.TextNode,
			Data: content,
		}})
}`, "go")
	d.AddSection("Charts and Images")
	d.AddImgFromURI("http://cs.smith.edu/~orourke/MathOverflow/GraphRand50.jpg", 3)
	d.AddChartLineXsYs([]float64{1, 2, 3, 4, 5, 6}, []float64{1.24, 2.123, 3.323, 3.9878, 4.143, 5.123}, 6, 3)

	d.AddSection("Table")
	d.AddParagraph("Tables are a nice way to organize a lot of data. We provide a few utility classes to help you style your table as easily as possible. In addition, to improve mobile experience, all tables on mobile-screen widths are centered automatically.")
	d.AddTable(
		[]string{"Name", "Item Name", "Item Price"},
		[][]string{
			{"Alvin", "Eclair", "$0.87"},
			{"Alan", "Jellybean", "$3.76"},
			{"Jonathan", "Lollipop", "$7.00"}})
	d.SaveTo("index.html")
}
