//Package mark provides a html mark element.
package mark

import (
	"github.com/qlova/seed"
)

//Seed is a mark.
type Seed struct {
	seed.Seed
}

//New returns a new mark.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("mark")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}