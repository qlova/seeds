//Package button provides a html button element.
package button

import (
	"github.com/qlova/seed"
)

//Seed is a button.
type Seed struct {
	seed.Seed
}

//New returns a new button.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("button")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}