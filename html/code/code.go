//Package code provides a html code element.
package code

import (
	"github.com/qlova/seed"
)

//Seed is a code.
type Seed struct {
	seed.Seed
}

//New returns a new code.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("code")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}