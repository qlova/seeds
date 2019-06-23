//Provides a textbox that hides its input with dots.
package passwordbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seeds/textbox"
)

type Seed struct {
	textbox.Seed
}

func New() Seed {
	var PasswordBox = textbox.New()

	PasswordBox.SetAttributes("type='password'")

	return Seed{PasswordBox}
}

func AddTo(parent seed.Interface) Seed {
	var PasswordBox = New()
	parent.Root().Add(PasswordBox)
	return PasswordBox
}
