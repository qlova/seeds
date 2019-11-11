//Package head provides a head element.
package head

import (
	"github.com/qlova/seed"
)

//Seed is a head.
type Seed struct {
	seed.Seed
}

//New returns a new head.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("head")
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}
