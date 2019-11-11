//Package textarea provides a html textarea element.
package textarea

import (
	"github.com/qlova/seed"
)

//Seed is a textarea.
type Seed struct {
	seed.Seed
}

//New returns a new textarea.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("textarea")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}