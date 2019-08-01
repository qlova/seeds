//Provides basic textbox for users to edit single-line text.
package textbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/script/global"
	"github.com/qlova/seed/style/css"
)

//Seed is a textbox that allows single-line text input.
type Seed struct {
	seed.Seed
	Fullscreen seed.Seed
}

//KeyboardVisible is a state that is set whenever an on-screen keyboard is detected.
var KeyboardVisible = seed.NewState()

//KeyboardHidden is a state that is set whenever an on-screen keyboard leaves.
var KeyboardHidden = KeyboardVisible.Not()

//New creates a new textbox.
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

	var save = global.NewString()

	var FullscreenTextBox = seed.AddTo(FullscreenEditor)
	FullscreenTextBox.SetTag("input")
	FullscreenTextBox.SetSize(100, 100)
	FullscreenTextBox.SetTextAlign(css.Center)

	FullscreenTextBox.OnEnter(func(q seed.Script) {
		KeyboardHidden.Set(q)
	})
	FullscreenTextBox.OnFocusLost(func(q seed.Script) {
		KeyboardHidden.Set(q)
	})
	FullscreenTextBox.OnClick(func(q seed.Script) {
		KeyboardHidden.Set(q)
	})

	TextBox.When(KeyboardVisible, func(q seed.Script) {
		FullscreenEditor.Script(q).SetVisible()
		FullscreenTextBox.Script(q).Focus()
		FullscreenTextBox.Script(q).SetValue(TextBox.Script(q).Value())
	})

	TextBox.When(KeyboardHidden, func(q seed.Script) {
		FullscreenTextBox.Script(q).Blur()
		TextBox.Script(q).SetValue(FullscreenTextBox.Script(q).Value())
		save.Set(q, TextBox.Script(q).Value())
		q.After(250, func() {
			FullscreenEditor.Script(q).SetHidden()
		})
	})

	TextBox.OnFocus(func(q seed.Script) {
		q.Javascript(`let height = document.body.clientHeight;`)
		q.After(100, func() {
			q.Javascript(`let done = false; for (t of [100, 250, 500, 600, 700, 800, 900, 1000]) { setTimeout(function(){if (!done && height > document.body.clientHeight) {done=true;`)

			KeyboardVisible.Set(q)

			q.Javascript(`let new_height = document.body.clientHeight; let f = function() {
				if (!done) return;
				
				if (document.body.clientHeight > new_height) {`)

			KeyboardHidden.Set(q)
			FullscreenEditor.Script(q).SetHidden()

			q.Javascript(`return;}
				
				setTimeout(f, 250);
			}; setTimeout(f, 1700);`)

			q.Javascript(`}}, t)}`)

		})
	})

	TextBox.OnChange(func(q seed.Script) {
		save.Set(q, TextBox.Script(q).Value())
	})
	TextBox.OnReady(func(q seed.Script) {
		TextBox.Script(q).SetValue(save.Get(q))
	})

	return Seed{TextBox, FullscreenTextBox}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var TextBox = New()
	parent.Root().Add(TextBox)
	return TextBox
}

//SetAttributes sets the HTML attributes for this seed.
func (textbox Seed) SetAttributes(attr string) {
	textbox.Seed.SetAttributes(attr)
	textbox.Fullscreen.SetAttributes(attr)
}

//SetPlaceholder sets the placeholder to appear when there is no usertext in the textbox.
func (textbox Seed) SetPlaceholder(placeholder string) {
	textbox.SetAttributes(textbox.Attributes() + " placeholder='" + placeholder + "' ")
}

//SetRequired sets this seed to be required within a parent form seed.
func (textbox Seed) SetRequired() {
	textbox.SetAttributes(textbox.Attributes() + " required")
}

//Script is the script context to this seed.
type Script struct {
	script.Seed
	Fullscreen script.Seed
}

//Script returns the script context to this seed.
func (textbox Seed) Script(q script.Script) Script {
	return Script{textbox.Seed.Script(q), textbox.Fullscreen.Script(q)}
}

//SetReadOnly disables the textbox from being edited.
func (textbox Script) SetReadOnly(state script.Bool) {
	textbox.Q.Javascript(textbox.Element() + ".readOnly = " + state.LanguageType().Raw() + ";")
}

//SetPlaceholder sets this seed's placeholder.
func (textbox Script) SetPlaceholder(placeholder script.String) {
	textbox.Q.Javascript(textbox.Element() + ".placeholder = " + placeholder.LanguageType().Raw() + ";")
	textbox.Q.Javascript(textbox.Fullscreen.Element() + ".placeholder = " + placeholder.LanguageType().Raw() + ";")
}
