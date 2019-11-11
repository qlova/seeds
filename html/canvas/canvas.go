//Package canvas provides a html canvas element.
package canvas

import (
	"github.com/qlova/seed"
)

//Seed is a canvas.
type Seed struct {
	seed.Seed
}

//New returns a new canvas.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("canvas")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}