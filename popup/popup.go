//Provides a simple popup template that will overlay the app. Hidden by default.
package popup

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seeds/column"
)

type Seed struct {
	column.Seed
}

func New() Seed {
	var Popup = column.New()

	Popup.SetPosition(css.Fixed)
	Popup.SetLeft(css.Decode(50))
	Popup.SetTop(css.Decode(50))
	Popup.Translate(-50, -50)
	Popup.Set("box-shadow", "3px 4px 20px black")

	Popup.SetSize(seed.Auto, seed.Auto)
	Popup.SetHidden()

	return Seed{Popup}
}

func AddTo(parent seed.Interface) Seed {
	var Popup = New()
	parent.Root().Add(Popup)
	return Popup
}
