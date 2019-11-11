//Package h4 provides a html h4 element.
package h4

import (
	"github.com/qlova/seed"
)

//Seed is a h4.
type Seed struct {
	seed.Seed
}

//New returns a new h4.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("h4")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}