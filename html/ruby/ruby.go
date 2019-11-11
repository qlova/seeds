//Package ruby provides a html ruby element.
package ruby

import (
	"github.com/qlova/seed"
)

//Seed is a ruby.
type Seed struct {
	seed.Seed
}

//New returns a new ruby.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("ruby")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}