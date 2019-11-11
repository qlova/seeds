//Package menu provides a html menu element.
package menu

import (
	"github.com/qlova/seed"
)

//Seed is a menu.
type Seed struct {
	seed.Seed
}

//New returns a new menu.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("menu")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}