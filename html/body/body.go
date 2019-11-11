//Package body provides a body element.
package body

import (
	"github.com/qlova/seed"
)

//Seed is a body.
type Seed struct {
	seed.Seed
}

//New returns a new body.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("body")
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}
