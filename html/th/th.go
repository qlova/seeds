//Package th provides a html th element.
package th

import (
	"github.com/qlova/seed"
)

//Seed is a th.
type Seed struct {
	seed.Seed
}

//New returns a new th.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("th")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}