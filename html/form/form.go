//Package form provides a html form element.
package form

import (
	"github.com/qlova/seed"
)

//Seed is a form.
type Seed struct {
	seed.Seed
}

//New returns a new form.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("form")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}