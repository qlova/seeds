//Provides basic display for text.
package text

import (
	"image/color"

	"github.com/qlova/seed"
	"github.com/qlova/seed/style"
)

type Seed struct {
	seed.Seed
}

//Set the text color.
func (text Seed) SetColor(c color.Color) {
	text.SetTextColor(c)
}

//Set the text's font-size.
func (text Seed) SetSize(s complex128) {
	text.SetTextSize(s)
}

//SetShadow sets the text's shadow.
func (text Seed) SetShadow(shadow style.Shadow) {
	text.SetTextShadow(shadow)
}

//Set the text's font-size.
func (text Seed) Underline() {
	text.Set("text-decoration", "underline")
}

func New(message ...string) Seed {
	var Text = seed.New()
	Text.SetTag("span")

	Text.ReactNative().SetTag("Text")

	if len(message) > 0 {
		Text.SetText(message[0])
		Text.ReactNative().SetContent("Text")
	}

	Text.SetSize(seed.Auto, seed.Auto)

	Text.Align(0)

	return Seed{Text}
}

func AddTo(parent seed.Interface, message ...string) Seed {
	var Text = New(message...)
	parent.Root().Add(Text)
	return Text
}
