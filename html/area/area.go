//Package area provides a html area element.
package area

import (
	"github.com/qlova/seed"
)

//Seed is a area.
type Seed struct {
	seed.Seed
}

//New returns a new area.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("area")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}