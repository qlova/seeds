//Package embed provides a html embed element.
package embed

import (
	"github.com/qlova/seed"
)

//Seed is a embed.
type Seed struct {
	seed.Seed
}

//New returns a new embed.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("embed")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}