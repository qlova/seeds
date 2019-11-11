//Package img provides a html img element.
package img

import (
	"github.com/qlova/seed"
)

//Seed is a img.
type Seed struct {
	seed.Seed
}

//New returns a new img.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("img")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}