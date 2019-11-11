//Package keygen provides a html keygen element.
package keygen

import (
	"github.com/qlova/seed"
)

//Seed is a keygen.
type Seed struct {
	seed.Seed
}

//New returns a new keygen.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("keygen")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}