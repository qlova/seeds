//Package htmlselect provides a html select element.
package htmlselect

import (
	"github.com/qlova/seed"
)

//Seed is a select.
type Seed struct {
	seed.Seed
}

//New returns a new select.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("select")
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}
