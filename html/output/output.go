//Package output provides a html output element.
package output

import (
	"github.com/qlova/seed"
)

//Seed is a output.
type Seed struct {
	seed.Seed
}

//New returns a new output.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("output")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}