//Package param provides a html param element.
package param

import (
	"github.com/qlova/seed"
)

//Seed is a param.
type Seed struct {
	seed.Seed
}

//New returns a new param.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("param")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}