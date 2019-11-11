//Package samp provides a html samp element.
package samp

import (
	"github.com/qlova/seed"
)

//Seed is a samp.
type Seed struct {
	seed.Seed
}

//New returns a new samp.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("samp")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}