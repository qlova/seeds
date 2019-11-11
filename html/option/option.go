//Package option provides a html option element.
package option

import (
	"github.com/qlova/seed"
)

//Seed is a option.
type Seed struct {
	seed.Seed
}

//New returns a new option.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("option")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}