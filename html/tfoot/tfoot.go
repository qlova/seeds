//Package tfoot provides a html tfoot element.
package tfoot

import (
	"github.com/qlova/seed"
)

//Seed is a tfoot.
type Seed struct {
	seed.Seed
}

//New returns a new tfoot.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("tfoot")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}