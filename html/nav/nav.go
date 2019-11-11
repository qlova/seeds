//Package nav provides a html nav element.
package nav

import (
	"github.com/qlova/seed"
)

//Seed is a nav.
type Seed struct {
	seed.Seed
}

//New returns a new nav.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("nav")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}