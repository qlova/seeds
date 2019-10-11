//Package p provides an html p element.
package p

import (
	"github.com/qlova/seed"
)

//Seed is a p.
type Seed struct {
	seed.Seed
}

//New returns a new p.
func New(message ...string) Seed {
	var Text = seed.New()
	Text.SetTag("p")

	if len(message) > 0 {
		Text.SetText(message[0])
	}

	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface, message ...string) Seed {
	var Text = New(message...)
	parent.Root().Add(Text)
	return Text
}
