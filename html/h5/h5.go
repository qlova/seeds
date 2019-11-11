//Package h5 provides a html h5 element.
package h5

import (
	"github.com/qlova/seed"
)

//Seed is a h5.
type Seed struct {
	seed.Seed
}

//New returns a new h5.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("h5")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}