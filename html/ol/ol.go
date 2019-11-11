//Package ol provides a html ol element.
package ol

import (
	"github.com/qlova/seed"
)

//Seed is a ol.
type Seed struct {
	seed.Seed
}

//New returns a new ol.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("ol")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}