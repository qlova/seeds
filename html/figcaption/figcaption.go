//Package figcaption provides a html figcaption element.
package figcaption

import (
	"github.com/qlova/seed"
)

//Seed is a figcaption.
type Seed struct {
	seed.Seed
}

//New returns a new figcaption.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("figcaption")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}