//Package h3 provides a html h3 element.
package h3

import (
	"github.com/qlova/seed"
)

//Seed is a h3.
type Seed struct {
	seed.Seed
}

//New returns a new h3.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("h3")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}