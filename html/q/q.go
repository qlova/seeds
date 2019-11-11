//Package q provides a html q element.
package q

import (
	"github.com/qlova/seed"
)

//Seed is a q.
type Seed struct {
	seed.Seed
}

//New returns a new q.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("q")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}