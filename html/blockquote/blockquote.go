//Package blockquote provides a html blockquote element.
package blockquote

import (
	"github.com/qlova/seed"
)

//Seed is a blockquote.
type Seed struct {
	seed.Seed
}

//New returns a new blockquote.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("blockquote")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}