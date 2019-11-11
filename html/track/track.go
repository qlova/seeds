//Package track provides a html track element.
package track

import (
	"github.com/qlova/seed"
)

//Seed is a track.
type Seed struct {
	seed.Seed
}

//New returns a new track.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("track")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}