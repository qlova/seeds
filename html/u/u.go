//Package u provides a html u element.
package u

import (
	"github.com/qlova/seed"
)

//Seed is a u.
type Seed struct {
	seed.Seed
}

//New returns a new u.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("u")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}