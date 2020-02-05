//Provides a button that can be clicked on to trigger an action.
package button

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/unit"
	"github.com/qlova/seeds/text"
)

type Seed struct {
	seed.Seed
	Text text.Seed
}

func New(label ...string) Seed {
	var Button = seed.New()

	Button.SetTag("button")

	//Fix issues with form validation.
	Button.Set("type", "button")

	Button.CSS().SetDisplay(css.Flex)

	Button.SetSize(unit.Auto, unit.Auto)
	Button.ItemSpacing().Outside()

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

type Ctx struct {
	script.Seed
	Text script.Seed
}

func (button Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{
		button.Seed.Ctx(q),
		button.Text.Ctx(q),
	}
}

func (button Ctx) SetText(text script.String) {
	button.Text.SetText(text)
}
