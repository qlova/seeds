//Package command provides a html command element.
package command

import (
	"github.com/qlova/seed"
)

//Seed is a command.
type Seed struct {
	seed.Seed
}

//New returns a new command.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("command")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}