//Package address provides a html address element.
package address

import (
	"github.com/qlova/seed"
)

//Seed is a address.
type Seed struct {
	seed.Seed
}

//New returns a new address.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("address")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}