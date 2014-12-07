// +build ignore
package main

import "github.com/gizak/docor"

func main() {
	d := docor.NewBlankDoc()

	d.AddTitle("Go Docor")
	d.AddBlockquote("Keep it simple, stupid")

	d.AddParagraph("Docor is simple, using that tex-thing like semantics to write report; Docor is stupid, only do one thing: generate one-page offline html report with 100% Go. Quick view:")

	d.AddCode(`    d := docor.NewBlankDoc()

	d.AddTitle("Go Docor")
	d.AddBlockquote("Keep it simple, stupid")

	d.AddParagraph("Docor is simple, using that tex-thing like semantics to write report; Docor is stupid, only do one thing: generate one-page offline html report with 100% Go.")
    //...

    d.SaveTo("index.html")`, "go")

	d.AddParagraph("Just that simple! Actually the idea behind this is from a feature I wanted in GoConvey: http://goo.gl/WbH1Ub . There is a scenario: when doing stochastic-related programming, often time I need a graph or table generator to display performance data on-the-fly and further more, generate a complete running report with pure Go programmingly. However, unfortunately there is no such thing that could meet my requirement, so I decided to make a super simple and stupid report generator with 100% Go!")

	d.AddSection("Installation")

	d.AddCode(`go get github.com/gizak/docor`, "bash")

	d.AddSection("Usage")

	d.AddParagraph("Docor is WIP, a basic set of writing API can be used, which contains image, drawing line graph, code snippets, paragraph, pre block, blockquote and etc.. more functions are comming!")

	d.AddSubsection("Image")

	d.AddCode(`//Add a simple image from url. url can be http or local address or even base64 src. An width needed, ranging from 1 to 12, int value.
d.AddImgFromURI(url string, width int)

//Add a detialed image
d.AddDetialImgFromURI(url string, width int, title, titleColor, description string)`, "go")

	d.AddImgFromURI("http://podcollective.com/portal/wp-content/uploads/2013/05/POLAR-INNER_LIGHT.jpg", 5)
	d.AddDetailImgFromURI("http://podcollective.com/portal/wp-content/uploads/2013/05/POLAR-SPIRIT_ORANGE.jpg", 5, "Polar Graph Art", "white", "VAJRA EQUATION from http://podcollective.com/portal/category/design/")

	d.AddSubsection("Typography")
	d.AddCode(`//Add page title, should appear on the top of your code
d.AddTitle(title string)

//Add section
d.AddSection(title string)

//Add Subsection
d.AddSubsection(title string)

//Add a blockquote
d.AddBlockquote(content string)

//Add a paragraph
d.AddParagraph(par string)`, "go")

	d.AddSubsection("Code")
	d.AddParagraph("Docor uses prism to render code highlight, nearly support all kinds of normal programming langauge. All code snippets in this page are rendered by this function.")
	d.AddCode(`//Add a code snippet
d.AddCode(code, lang string)`, "go")

	d.AddSubsection("Chart")
	d.AddParagraph("Docor uses plotinum as drawing library. At this moment Docor only has the line chart API, but more and more APIs are on their way!")
	d.AddChartLineXsYs([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{1.24, 2.123, 3.323, 3.9878, 4.143, 5.123, 6.66, 6.99, 7.32}, 5, 3)
	d.AddCode(`//Add a line chart by x axis points and y axis points.
//One should also assign width and height to chart.
//Both width and height should between 1 and 12 inclusively
d.AddChartLineXsYs(xs, ys []float64, width, height int)

// e.g.
d.AddChartLineXsYs([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{1.24, 2.123, 3.323, 3.9878, 4.143, 5.123, 6.66, 6.99, 7.32}, 5, 3)`, "go")

	d.AddSubsection("Table")
	d.AddParagraph("A table helps dump the data!")
	d.AddTable(
		[]string{"Name", "Item Name", "Item Price", "Payment"},
		[][]string{
			{"Alvin", "Eclair", "$0.87", "Cash"},
			{"Alan", "Jellybean", "$3.76", "Cash"},
			{"Jonathan", "Lollipop", "$7.00", "Credit"}})
	d.AddCode(`//Add a table. A table need a table head
//table only accept string value
d.AddTable(thead []string, tbody[][]string)
//e.g.
d.AddTable(
		[]string{"Name", "Item Name", "Item Price", "Payment"},
		[][]string{
			{"Alvin", "Eclair", "$0.87", "Cash"},
			{"Alan", "Jellybean", "$3.76", "Cash"},
			{"Jonathan", "Lollipop", "$7.00", "Credit"}})
`, "go")

	d.AddSubsection("Pre Block")
	d.AddParagraph("A pre block is nothing but a html <pre> block, just like following: (I am considering if it needs a frame)")
	d.AddPre(`<!DOCTYPE html>
<html>
<head>
 <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" >
 <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" >

 <meta name="ROBOTS" content="NOARCHIVE">`)

	d.AddCode(`//A html pre block
d.AddPre(content string)`, "go")

	d.AddSection("Development")
	d.AddParagraph("Hmm, let me think...")
	d.SaveTo("index.html")
}
