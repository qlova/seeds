//Provides basic textbox for users to edit single-line text.
package textbox

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var TextBox = seed.New()

	TextBox.SetTag("input")

	TextBox.SetSize(seed.Auto, seed.Auto)

	TextBox.Align(0)

	var save = script.NewString()
	TextBox.OnChange(func(q seed.Script) {
		q.Set(save, TextBox.Script(q).Value())
	})
	TextBox.OnReady(func(q seed.Script) {
		TextBox.Script(q).SetValue(save.Script(q))
	})

	return Seed{TextBox}
}

func AddTo(parent seed.Interface) Seed {
	var TextBox = New()
	parent.Root().Add(TextBox)
	return TextBox
}

func (textbox Seed) SetRequired() {
	textbox.SetAttributes(textbox.Attributes() + " required")
}

type Script struct {
	script.Seed
}

func (textbox Seed) Script(q script.Script) Script {
	return Script{textbox.Seed.Script(q)}
}

func (textbox Script) SetReadOnly(state script.Bool) {
	textbox.Q.Javascript(textbox.Element() + ".readOnly = " + state.LanguageType().Raw() + ";")
}
