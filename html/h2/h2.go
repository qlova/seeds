//Package h2 provides a html h2 element.
package h2

import (
	"github.com/qlova/seed"
)

//Seed is a h2.
type Seed struct {
	seed.Seed
}

//New returns a new h2.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("h2")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}