//Package bdo provides a html bdo element.
package bdo

import (
	"github.com/qlova/seed"
)

//Seed is a bdo.
type Seed struct {
	seed.Seed
}

//New returns a new bdo.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("bdo")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}