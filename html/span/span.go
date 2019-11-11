//Package span provides a html span element.
package span

import (
	"github.com/qlova/seed"
)

//Seed is a span.
type Seed struct {
	seed.Seed
}

//New returns a new span.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("span")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}