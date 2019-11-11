package main

import (
	"fmt"
	"os"
)

var Elements = []string{
	"a",
	"abbr",
	"address",
	"area",
	"article",
	"aside",
	"audio",
	"b",
	"bdi",
	"bdo",
	"blockquote",
	"body",
	"br",
	"button",
	"canvas",
	"caption",
	"cite",
	"code",
	"col",
	"colgroup",
	"command",
	"datalist",
	"dd",
	"del",
	"details",
	"dfn",
	"div",
	"dl",
	"dt",
	"em",
	"embed",
	"fieldset",
	"figcaption",
	"figure",
	"footer",
	"form",
	"h1",
	"h2",
	"h3",
	"h4",
	"h5",
	"h6",
	"header",
	"hr",
	"html",
	"i",
	"iframe",
	"img",
	"input",
	"ins",
	"kbd",
	"keygen",
	"label",
	"legend",
	"li",
	"main",
	"map",
	"mark",
	"menu",
	"meter",
	"nav",
	"object",
	"ol",
	"optgroup",
	"option",
	"output",
	"p",
	"param",
	"pre",
	"progress",
	"q",
	"rp",
	"rt",
	"ruby",
	"s",
	"samp",
	"section",
	"select",
	"small",
	"source",
	"span",
	"strong",
	"sub",
	"summary",
	"sup",
	"table",
	"tbody",
	"td",
	"textarea",
	"tfoot",
	"th",
	"thead",
	"time",
	"tr",
	"track",
	"u",
	"ul",
	"var",
	"video",
	"wbr",
}

func main() {
	const HTMLPackagePath = "../../"

	for _, element := range Elements {
		os.Mkdir(HTMLPackagePath+element, 0755)

		if fi, err := os.Open(HTMLPackagePath + element + "/" + element + ".go"); err == nil {
			fi.Close()
			continue
		}

		if fi, err := os.Create(HTMLPackagePath + element + "/" + element + ".go"); err == nil {
			switch element {
			case "select", "var", "map", "main":
				element = "html" + element
			}

			fmt.Fprintf(fi, `//Package %[1]v provides a html %[1]v element.
package %[1]v

import (
	"github.com/qlova/seed"
)

//Seed is a %[1]v.
type Seed struct {
	seed.Seed
}

//New returns a new %[1]v.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("%[1]v")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}`, element)
			fi.Close()
		}
	}
}
