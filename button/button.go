//Provides a button that can be clicked on to trigger an action.
package button

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"
import "github.com/qlova/seeds/text"

type Seed struct {
	seed.Seed
	Text text.Seed
}

func New(label ...string) Seed {
	var Button = seed.New()

	Button.SetTag("button")

	//Fix issues with form validation.
	Button.Set("type", "button")

	Button.SetSize(seed.Auto, seed.Auto)
	Button.AlignChildren(0)

	var Text = text.AddTo(Button, label...)

	return Seed{Button, Text}
}

func AddTo(parent seed.Interface, label ...string) Seed {
	var Button = New(label...)
	parent.Root().Add(Button)
	return Button
}

func (button Seed) SetText(text string) {
	button.Text.SetText(text)
}

type Script struct {
	script.Seed
	Text script.Seed
}

func (button Seed) Script(q seed.Script) Script {
	return Script{
		button.Seed.Script(q),
		button.Text.Script(q),
	}
}

func (button Script) SetText(text script.String) {
	button.Text.SetText(text)
}
