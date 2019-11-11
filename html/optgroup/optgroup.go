//Package optgroup provides a html optgroup element.
package optgroup

import (
	"github.com/qlova/seed"
)

//Seed is a optgroup.
type Seed struct {
	seed.Seed
}

//New returns a new optgroup.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("optgroup")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}