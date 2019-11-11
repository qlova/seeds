//Package dfn provides a html dfn element.
package dfn

import (
	"github.com/qlova/seed"
)

//Seed is a dfn.
type Seed struct {
	seed.Seed
}

//New returns a new dfn.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("dfn")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}