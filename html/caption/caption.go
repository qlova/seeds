//Package caption provides a html caption element.
package caption

import (
	"github.com/qlova/seed"
)

//Seed is a caption.
type Seed struct {
	seed.Seed
}

//New returns a new caption.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("caption")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}