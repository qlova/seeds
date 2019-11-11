//Package datalist provides a html datalist element.
package datalist

import (
	"github.com/qlova/seed"
)

//Seed is a datalist.
type Seed struct {
	seed.Seed
}

//New returns a new datalist.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("datalist")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}