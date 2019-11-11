//Package pre provides a html pre element.
package pre

import (
	"github.com/qlova/seed"
)

//Seed is a pre.
type Seed struct {
	seed.Seed
}

//New returns a new pre.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("pre")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}