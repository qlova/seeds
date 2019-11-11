//Package tr provides a html tr element.
package tr

import (
	"github.com/qlova/seed"
)

//Seed is a tr.
type Seed struct {
	seed.Seed
}

//New returns a new tr.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("tr")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}