//Package popup provides a simple popup template that will overlay the app. Hidden by default.
package popup

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
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

	Popup.ItemSpacing().Center()
	Popup.AlignItems().Center()
	Popup.SetHidden()
	Popup.SetColor(seed.RGBA(0, 0, 0, 128))
	Popup.SetLayer(999)
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

func (popup Seed) SetTransition(trans seed.Transition) {
	if trans.In != nil || !trans.When.Null() || trans.WhenTag != "" || trans.WhenBack {
		popup.OnShow(func(q script.Ctx) {
			var Page = popup.Ctx(q)
			seed.SetTransitionIn(Page, trans)
		})
	}
	if trans.Out != nil || !trans.When.Null() || trans.WhenTag != "" || trans.WhenBack {
		popup.OnHide(func(q script.Ctx) {
			var Page = popup.Ctx(q)
			seed.SetTransitionOut(Page, trans)
		})
	}
}

const js = `seeds.popup = {};
seeds.popup.show = async function(element) {
	if (getComputedStyle(element).display == "none") {
		element.style.display = "flex";
		if (element.onshow) {
			await element.onshow();
			if (seed.goto.in) {
				await seed.goto.in;
				seed.goto.in = null;
			}
		}
	}
};

seeds.popup.hide = async function(element) {
	if (getComputedStyle(element).display != "none") {
		element.style.display = "none";
		if (element.onhide) {
			await element.onhide();
			if (seed.goto.out) {
				await seed.goto.out;
				seed.goto.out = null;
			}
		}
	}
};
`

//Show this popup.
func (popup Seed) Show(q script.Ctx) {
	q.Require(js)
	q.Javascript("seeds.popup.show(%v);", popup.Ctx(q).Element())
}

//OnShow runs the callback when the popup is shown.
func (popup Seed) OnShow(s func(q script.Ctx)) {
	popup.On("show", s)
}

//Hide this popup.
func (popup Seed) Hide(q script.Ctx) {
	q.Require(js)
	q.Javascript("seeds.popup.hide(%v);", popup.Ctx(q).Element())
}

//OnHide runs the callback when the popup is shown.
func (popup Seed) OnHide(s func(q script.Ctx)) {
	popup.On("hide", s)
}
