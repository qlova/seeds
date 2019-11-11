//Package dd provides a html dd element.
package dd

import (
	"github.com/qlova/seed"
)

//Seed is a dd.
type Seed struct {
	seed.Seed
}

//New returns a new dd.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("dd")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}