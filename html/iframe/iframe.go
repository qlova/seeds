//Package iframe provides a html iframe element.
package iframe

import (
	"github.com/qlova/seed"
)

//Seed is a iframe.
type Seed struct {
	seed.Seed
}

//New returns a new iframe.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("iframe")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}