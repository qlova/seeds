//Package wbr provides a html wbr element.
package wbr

import (
	"github.com/qlova/seed"
)

//Seed is a wbr.
type Seed struct {
	seed.Seed
}

//New returns a new wbr.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("wbr")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}