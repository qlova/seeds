//Package summary provides a html summary element.
package summary

import (
	"github.com/qlova/seed"
)

//Seed is a summary.
type Seed struct {
	seed.Seed
}

//New returns a new summary.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("summary")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}