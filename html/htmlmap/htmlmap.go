//Package htmlmap provides a html map element.
package htmlmap

import (
	"github.com/qlova/seed"
)

//Seed is a map.
type Seed struct {
	seed.Seed
}

//New returns a new map.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("map")
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}
