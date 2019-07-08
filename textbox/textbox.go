//Provides basic textbox for users to edit single-line text.
package textbox

import "github.com/qlova/seed"
import "github.com/qlova/seed/style/css"
import "github.com/qlova/seed/script"

type Seed struct {
	seed.Seed
	Fullscreen seed.Seed
}

func New() Seed {
	var TextBox = seed.New()

	TextBox.SetTag("input")

	TextBox.SetSize(seed.Auto, seed.Auto)

	TextBox.Align(0)

	var FullscreenEditor = seed.AddTo(TextBox)
	FullscreenEditor.SetHidden()
	FullscreenEditor.SetColor(seed.White)
	FullscreenEditor.SetLayer(99999)
	FullscreenEditor.CSS().SetPosition(css.Fixed)
	FullscreenEditor.CSS().SetTop(css.Zero)
	FullscreenEditor.CSS().SetLeft(css.Zero)
	FullscreenEditor.CSS().SetWidth(css.Number(100).Vw())
	FullscreenEditor.CSS().SetHeight(css.Number(100).Vh())

	var save = script.NewString()

	var FullscreenTextBox = seed.AddTo(FullscreenEditor)
	FullscreenTextBox.SetTag("input")
	FullscreenTextBox.SetSize(100, 100)
	FullscreenTextBox.SetTextAlign(css.Center)

	var f = seed.NewFunction(func(q seed.Script) {
		FullscreenTextBox.Script(q).Blur()
		TextBox.Script(q).SetValue(FullscreenTextBox.Script(q).Value())
		q.Set(save, TextBox.Script(q).Value())
		q.After(250, func() {
			FullscreenEditor.Script(q).SetHidden()
		})
	})

	FullscreenTextBox.OnEnter(func(q seed.Script) {
		q.Run(f)
	})
	FullscreenTextBox.OnFocusLost(func(q seed.Script) {
		q.Run(f)
	})
	FullscreenTextBox.OnClick(func(q seed.Script) {
		q.Run(f)
	})

	TextBox.OnFocus(func(q seed.Script) {
		q.Javascript(`let height = document.body.clientHeight;`)
		q.After(100, func() {
			q.Javascript(`let done = false; for (t of [100, 250, 500, 600, 700, 800, 900, 1000]) { setTimeout(function(){if (!done && height > document.body.clientHeight) {done=true;`)

			FullscreenEditor.Script(q).SetVisible()
			FullscreenTextBox.Script(q).Focus()
			FullscreenTextBox.Script(q).SetValue(TextBox.Script(q).Value())

			q.Javascript(`let new_height = document.body.clientHeight; let f = function() {
				if (!done) return;
				
				if (document.body.clientHeight > new_height) {`)

			FullscreenTextBox.Script(q).Blur()
			TextBox.Script(q).SetValue(FullscreenTextBox.Script(q).Value())
			q.Set(save, TextBox.Script(q).Value())

			FullscreenEditor.Script(q).SetHidden()

			q.Javascript(`return;}
				
				setTimeout(f, 250);
			}; setTimeout(f, 1700);`)

			q.Javascript(`}}, t)}`)

		})
	})

	TextBox.OnChange(func(q seed.Script) {
		q.Set(save, TextBox.Script(q).Value())
	})
	TextBox.OnReady(func(q seed.Script) {
		TextBox.Script(q).SetValue(save.Script(q))
	})

	return Seed{TextBox, FullscreenTextBox}
}

func AddTo(parent seed.Interface) Seed {
	var TextBox = New()
	parent.Root().Add(TextBox)
	return TextBox
}

func (textbox Seed) SetAttributes(attr string) {
	textbox.Seed.SetAttributes(attr)
	textbox.Fullscreen.SetAttributes(attr)
}

func (textbox Seed) SetPlaceholder(placeholder string) {
	textbox.SetAttributes(textbox.Attributes() + " placeholder='" + placeholder + "' ")
}

func (textbox Seed) SetRequired() {
	textbox.SetAttributes(textbox.Attributes() + " required")
}

type Script struct {
	script.Seed
	Fullscreen script.Seed
}

func (textbox Seed) Script(q script.Script) Script {
	return Script{textbox.Seed.Script(q), textbox.Fullscreen.Script(q)}
}

func (textbox Script) SetReadOnly(state script.Bool) {
	textbox.Q.Javascript(textbox.Element() + ".readOnly = " + state.LanguageType().Raw() + ";")
}

func (textbox Script) SetPlaceholder(placeholder script.String) {
	textbox.Q.Javascript(textbox.Element() + ".placeholder = " + placeholder.LanguageType().Raw() + ";")
	textbox.Q.Javascript(textbox.Fullscreen.Element() + ".placeholder = " + placeholder.LanguageType().Raw() + ";")
}
