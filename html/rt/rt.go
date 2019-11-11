//Package rt provides a html rt element.
package rt

import (
	"github.com/qlova/seed"
)

//Seed is a rt.
type Seed struct {
	seed.Seed
}

//New returns a new rt.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("rt")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}