//Package strong provides a html strong element.
package strong

import (
	"github.com/qlova/seed"
)

//Seed is a strong.
type Seed struct {
	seed.Seed
}

//New returns a new strong.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("strong")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}