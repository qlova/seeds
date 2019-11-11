//Package td provides a html td element.
package td

import (
	"github.com/qlova/seed"
)

//Seed is a td.
type Seed struct {
	seed.Seed
}

//New returns a new td.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("td")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}