//Package htmlmain provides a html main element.
package htmlmain

import (
	"github.com/qlova/seed"
)

//Seed is a main.
type Seed struct {
	seed.Seed
}

//New returns a new main.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("main")
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}
