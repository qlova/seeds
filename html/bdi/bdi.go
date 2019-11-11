//Package bdi provides a html bdi element.
package bdi

import (
	"github.com/qlova/seed"
)

//Seed is a bdi.
type Seed struct {
	seed.Seed
}

//New returns a new bdi.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("bdi")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}