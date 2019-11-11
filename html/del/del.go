//Package del provides a html del element.
package del

import (
	"github.com/qlova/seed"
)

//Seed is a del.
type Seed struct {
	seed.Seed
}

//New returns a new del.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("del")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}