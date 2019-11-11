//Package section provides a html section element.
package section

import (
	"github.com/qlova/seed"
)

//Seed is a section.
type Seed struct {
	seed.Seed
}

//New returns a new section.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("section")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}