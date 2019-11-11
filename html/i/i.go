//Package i provides a html i element.
package i

import (
	"github.com/qlova/seed"
)

//Seed is a i.
type Seed struct {
	seed.Seed
}

//New returns a new i.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("i")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}