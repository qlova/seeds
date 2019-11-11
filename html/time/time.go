//Package time provides a html time element.
package time

import (
	"github.com/qlova/seed"
)

//Seed is a time.
type Seed struct {
	seed.Seed
}

//New returns a new time.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("time")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}