//Package meter provides a html meter element.
package meter

import (
	"github.com/qlova/seed"
)

//Seed is a meter.
type Seed struct {
	seed.Seed
}

//New returns a new meter.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("meter")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}