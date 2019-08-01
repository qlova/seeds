//Package popup provides a simple popup template that will overlay the app. Hidden by default.
package popup

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seeds/column"
)

//Seed is the popup seed.
type Seed struct {
	column.Seed
}

//New returns a new popup.
func New() Seed {
	var Popup = column.New()

	Popup.AlignChildren(0)
	Popup.SetChildAlignment(0)
	Popup.SetHidden()
	Popup.SetColor(seed.RGBA(0, 0, 0, 128))
	Popup.SetLayer(999)
	Popup.CenterChildren()
	Popup.CSS().SetPosition(css.Fixed)
	Popup.CSS().SetTop(css.Zero)
	Popup.CSS().SetLeft(css.Zero)
	Popup.CSS().SetWidth(css.Number(100).Vw())
	Popup.CSS().SetHeight(css.Number(100).Vh())

	return Seed{Popup}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Popup = New()
	parent.Root().Add(Popup)
	return Popup
}

//Show this popup.
func (popup Seed) Show(q seed.Script) {
	popup.Script(q).SetVisible()
}

//Hide this popup.
func (popup Seed) Hide(q seed.Script) {
	popup.Script(q).SetHidden()
}
