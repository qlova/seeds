//Package em provides a html em element.
package em

import (
	"github.com/qlova/seed"
)

//Seed is a em.
type Seed struct {
	seed.Seed
}

//New returns a new em.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("em")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}