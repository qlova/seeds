//Package dl provides a html dl element.
package dl

import (
	"github.com/qlova/seed"
)

//Seed is a dl.
type Seed struct {
	seed.Seed
}

//New returns a new dl.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("dl")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}