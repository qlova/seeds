//Package footer provides a html footer element.
package footer

import (
	"github.com/qlova/seed"
)

//Seed is a footer.
type Seed struct {
	seed.Seed
}

//New returns a new footer.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("footer")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}