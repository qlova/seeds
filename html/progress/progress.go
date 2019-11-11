//Package progress provides a html progress element.
package progress

import (
	"github.com/qlova/seed"
)

//Seed is a progress.
type Seed struct {
	seed.Seed
}

//New returns a new progress.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("progress")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}