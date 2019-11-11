//Package dt provides a html dt element.
package dt

import (
	"github.com/qlova/seed"
)

//Seed is a dt.
type Seed struct {
	seed.Seed
}

//New returns a new dt.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("dt")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}