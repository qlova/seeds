//Package col provides a html col element.
package col

import (
	"github.com/qlova/seed"
)

//Seed is a col.
type Seed struct {
	seed.Seed
}

//New returns a new col.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("col")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}