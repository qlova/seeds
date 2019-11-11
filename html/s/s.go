//Package s provides a html s element.
package s

import (
	"github.com/qlova/seed"
)

//Seed is a s.
type Seed struct {
	seed.Seed
}

//New returns a new s.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("s")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}