//Package colgroup provides a html colgroup element.
package colgroup

import (
	"github.com/qlova/seed"
)

//Seed is a colgroup.
type Seed struct {
	seed.Seed
}

//New returns a new colgroup.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("colgroup")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}