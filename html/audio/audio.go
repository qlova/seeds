//Package audio provides a html audio element.
package audio

import (
	"github.com/qlova/seed"
)

//Seed is a audio.
type Seed struct {
	seed.Seed
}

//New returns a new audio.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("audio")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}