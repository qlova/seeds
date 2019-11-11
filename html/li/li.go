//Package li provides a html li element.
package li

import (
	"github.com/qlova/seed"
)

//Seed is a li.
type Seed struct {
	seed.Seed
}

//New returns a new li.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("li")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}