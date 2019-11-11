//Package object provides a html object element.
package object

import (
	"github.com/qlova/seed"
)

//Seed is a object.
type Seed struct {
	seed.Seed
}

//New returns a new object.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("object")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}