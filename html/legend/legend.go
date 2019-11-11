//Package legend provides a html legend element.
package legend

import (
	"github.com/qlova/seed"
)

//Seed is a legend.
type Seed struct {
	seed.Seed
}

//New returns a new legend.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("legend")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}