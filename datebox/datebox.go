//Provides a date-picker inputbox.
package datebox

import "github.com/qlova/seed"
import "github.com/qlova/seed/style/css"
import "github.com/qlova/seed/script"
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
	FakeBox.SetOpacity(css.Zero)
	FakeBox.SetPosition(css.Absolute)
	FakeBox.SetPointerEvents(css.None)
	
	FakeBox.OnChange(func(q seed.Script) {
		TextBox.Script(q).SetValue(FakeBox.Script(q).Value())
	})

	TextBox.OnChange(func(q seed.Script) {
		FakeBox.Script(q).SetValue(TextBox.Script(q).Value())
	})

	TextBox.OnClick(func(q seed.Script) {
		q.Javascript(`if (/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {`)
		FakeBox.Script(q).Focus()
		FakeBox.Script(q).Click()
		q.Javascript(`}`)
	})

	TextBox.OnReady(func(q seed.Script) {
		q.Javascript(`if (!(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent))) {`)
		q.Javascript(TextBox.Script(q).Element() + `.readOnly = false;`)
		q.Javascript(TextBox.Script(q).Element() + `.type = "date";`)
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

func (datebox Seed) Script(q seed.Script) Script {
	return Script{
		datebox.Seed.Script(q),
		datebox.fakebox.Seed.Script(q),
	}
}

type Script struct {
	script.Seed
	fakebox script.Seed
}

func (script Script) SetValue(value qlova.String) {
	script.Seed.SetValue(value)
	script.fakebox.SetValue(value)
}
