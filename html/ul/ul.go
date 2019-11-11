//Package ul provides a html ul element.
package ul

import (
	"github.com/qlova/seed"
)

//Seed is a ul.
type Seed struct {
	seed.Seed
}

//New returns a new ul.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("ul")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}