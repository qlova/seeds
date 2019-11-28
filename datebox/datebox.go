//Provides a date-picker inputbox.
package datebox

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"
import "github.com/qlova/seed/style/css"
import "github.com/qlova/seeds/textbox"

type Seed struct {
	textbox.Seed
	fakebox textbox.Seed
}

func New() Seed {
	var TextBox = textbox.New()
	TextBox.SetAttributes("readonly")

	var FakeBox = textbox.AddTo(TextBox)

	FakeBox.SetAttributes("type='date'")
	FakeBox.SetOpacity(0)
	FakeBox.CSS().SetPosition(css.Absolute)
	FakeBox.CSS().SetPointerEvents(css.None)

	FakeBox.OnChange(func(q script.Ctx) {
		TextBox.Ctx(q).SetValue(FakeBox.Ctx(q).Value())
	})

	TextBox.OnChange(func(q script.Ctx) {
		FakeBox.Ctx(q).SetValue(TextBox.Ctx(q).Value())
	})

	TextBox.OnClick(func(q script.Ctx) {
		q.Javascript(`if (/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {`)
		FakeBox.Ctx(q).Focus()
		FakeBox.Ctx(q).Click()
		q.Javascript(`}`)
	})

	TextBox.OnReady(func(q script.Ctx) {
		q.Javascript(`if (!(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent))) {`)
		q.Javascript(TextBox.Ctx(q).Element() + `.readOnly = false;`)
		q.Javascript(TextBox.Ctx(q).Element() + `.type = "date";`)
		q.Javascript(`}`)
	})

	return Seed{TextBox, FakeBox}
}

func (datebox Seed) SetRequired() {
	datebox.fakebox.SetRequired()
}

func AddTo(parent seed.Interface) Seed {
	var DateBox = New()
	parent.Root().Add(DateBox)
	return DateBox
}

func (datebox Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{
		datebox.Seed.Ctx(q),
		datebox.fakebox.Seed.Ctx(q),
	}
}

type Ctx struct {
	textbox.Ctx
	fakebox script.Seed
}

func (script Ctx) SetValue(value script.String) {
	script.Seed.SetValue(value)
	script.fakebox.SetValue(value)
}
