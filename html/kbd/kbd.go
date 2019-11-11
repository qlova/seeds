//Package kbd provides a html kbd element.
package kbd

import (
	"github.com/qlova/seed"
)

//Seed is a kbd.
type Seed struct {
	seed.Seed
}

//New returns a new kbd.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("kbd")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}