//Package ins provides a html ins element.
package ins

import (
	"github.com/qlova/seed"
)

//Seed is a ins.
type Seed struct {
	seed.Seed
}

//New returns a new ins.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("ins")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}