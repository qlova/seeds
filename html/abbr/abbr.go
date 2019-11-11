//Package abbr provides a html abbr element.
package abbr

import (
	"github.com/qlova/seed"
)

//Seed is a abbr.
type Seed struct {
	seed.Seed
}

//New returns a new abbr.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("abbr")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}