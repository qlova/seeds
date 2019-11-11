//Package source provides a html source element.
package source

import (
	"github.com/qlova/seed"
)

//Seed is a source.
type Seed struct {
	seed.Seed
}

//New returns a new source.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("source")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}