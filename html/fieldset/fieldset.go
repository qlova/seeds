//Package fieldset provides a html fieldset element.
package fieldset

import (
	"github.com/qlova/seed"
)

//Seed is a fieldset.
type Seed struct {
	seed.Seed
}

//New returns a new fieldset.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("fieldset")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}