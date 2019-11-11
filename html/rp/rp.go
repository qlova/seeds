//Package rp provides a html rp element.
package rp

import (
	"github.com/qlova/seed"
)

//Seed is a rp.
type Seed struct {
	seed.Seed
}

//New returns a new rp.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("rp")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}