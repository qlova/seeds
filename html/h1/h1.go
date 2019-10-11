//Package h1 provides an html h1 element.
package h1

import (
	"github.com/qlova/seed"
)

//Seed is a h1.
type Seed struct {
	seed.Seed
}

//New returns a new h1.
func New(message ...string) Seed {
	var Text = seed.New()
	Text.SetTag("h1")

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
