//Package style provides a style element.
package style

import (
	"github.com/qlova/seed"
)

//Seed is a style.
type Seed struct {
	seed.Seed
}

//New returns a new style.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("style")
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}
