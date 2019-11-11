//Package sup provides a html sup element.
package sup

import (
	"github.com/qlova/seed"
)

//Seed is a sup.
type Seed struct {
	seed.Seed
}

//New returns a new sup.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("sup")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}