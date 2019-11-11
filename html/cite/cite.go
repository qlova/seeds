//Package cite provides a html cite element.
package cite

import (
	"github.com/qlova/seed"
)

//Seed is a cite.
type Seed struct {
	seed.Seed
}

//New returns a new cite.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("cite")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}