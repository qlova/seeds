//Package div provides a html div element.
package div

import (
	"github.com/qlova/seed"
)

//Seed is a div.
type Seed struct {
	seed.Seed
}

//New returns a new div.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("div")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}