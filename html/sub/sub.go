//Package sub provides a html sub element.
package sub

import (
	"github.com/qlova/seed"
)

//Seed is a sub.
type Seed struct {
	seed.Seed
}

//New returns a new sub.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("sub")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}