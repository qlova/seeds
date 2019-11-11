//Package h6 provides a html h6 element.
package h6

import (
	"github.com/qlova/seed"
)

//Seed is a h6.
type Seed struct {
	seed.Seed
}

//New returns a new h6.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("h6")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}