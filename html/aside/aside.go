//Package aside provides a html aside element.
package aside

import (
	"github.com/qlova/seed"
)

//Seed is a aside.
type Seed struct {
	seed.Seed
}

//New returns a new aside.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("aside")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}