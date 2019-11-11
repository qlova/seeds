//Package b provides a html b element.
package b

import (
	"github.com/qlova/seed"
)

//Seed is a b.
type Seed struct {
	seed.Seed
}

//New returns a new b.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("b")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}