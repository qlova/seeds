//Package meta provides a meta element.
package meta

import (
	"github.com/qlova/seed"
)

//Seed is a meta.
type Seed struct {
	seed.Seed
}

//New returns a new meta.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("meta")
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}
