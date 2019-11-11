//Package title provides an html title element.
package title

import (
	"github.com/qlova/seed"
)

//Seed is a title.
type Seed struct {
	seed.Seed
}

//New returns a new title.
func New(message ...string) Seed {
	var Text = seed.New()
	Text.SetTag("title")

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
