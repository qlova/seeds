//Package link provides a link element.
package link

import (
	"github.com/qlova/seed"
)

//Seed is a link.
type Seed struct {
	seed.Seed
}

//New returns a new link.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("link")
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}
