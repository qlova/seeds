//Package details provides a html details element.
package details

import (
	"github.com/qlova/seed"
)

//Seed is a details.
type Seed struct {
	seed.Seed
}

//New returns a new details.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("details")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}