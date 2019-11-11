//Package html provides an html element.
package html

import (
	"github.com/qlova/seed"
)

//Seed is a html.
type Seed struct {
	seed.Seed
}

//New returns a new html.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("html")
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}
