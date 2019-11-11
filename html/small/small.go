//Package small provides a html small element.
package small

import (
	"github.com/qlova/seed"
)

//Seed is a small.
type Seed struct {
	seed.Seed
}

//New returns a new small.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("small")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}