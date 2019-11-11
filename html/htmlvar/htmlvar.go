//Package htmlvar provides a html var element.
package htmlvar

import (
	"github.com/qlova/seed"
)

//Seed is a var.
type Seed struct {
	seed.Seed
}

//New returns a new var.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("var")
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}
