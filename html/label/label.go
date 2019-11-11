//Package label provides a html label element.
package label

import (
	"github.com/qlova/seed"
)

//Seed is a label.
type Seed struct {
	seed.Seed
}

//New returns a new label.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("label")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}