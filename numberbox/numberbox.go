//Provides a textbox that only accepts numbers.
package numberbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seeds/textbox"
)

type Seed struct {
	textbox.Seed
}

func New() Seed {
	var NumberBox = textbox.New()

	NumberBox.SetAttributes(`inputmode="numeric" pattern="[0-9]*" type='number'`)
	NumberBox.Fullscreen.SetAttributes(`inputmode="numeric" pattern="[0-9]*" type='number'`)

	return Seed{NumberBox}
}

func AddTo(parent seed.Interface) Seed {
	var NumberBox = New()
	parent.Root().Add(NumberBox)
	return NumberBox
}
