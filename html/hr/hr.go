//Package hr provides a html hr element.
package hr

import (
	"github.com/qlova/seed"
)

//Seed is a hr.
type Seed struct {
	seed.Seed
}

//New returns a new hr.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("hr")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}