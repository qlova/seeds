//Package input provides a html input element.
package input

import (
	"github.com/qlova/seed"
)

//Seed is a input.
type Seed struct {
	seed.Seed
}

//New returns a new input.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("input")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}