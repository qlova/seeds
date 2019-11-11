//Package thead provides a html thead element.
package thead

import (
	"github.com/qlova/seed"
)

//Seed is a thead.
type Seed struct {
	seed.Seed
}

//New returns a new thead.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("thead")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}