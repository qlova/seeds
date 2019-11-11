//Package figure provides a html figure element.
package figure

import (
	"github.com/qlova/seed"
)

//Seed is a figure.
type Seed struct {
	seed.Seed
}

//New returns a new figure.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("figure")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}