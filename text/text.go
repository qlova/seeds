//Provides basic display for text.
package text

import (
	"image/color"

	"github.com/qlova/seed"
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/unit"
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

	Text.SetSize(unit.Auto, unit.Auto)

	Text.Align().Center()
	Text.TextAlign().Center()

	return Seed{Text}
}

func AddTo(parent seed.Interface, message ...string) Seed {
	var Text = New(message...)
	parent.Root().Add(Text)
	return Text
}

//Aligner aligns text.
type Aligner struct {
	style style.Style
}

//Left aligned text.
func (a Aligner) Left() {
	a.style.SetAlignSelf(css.FlexStart)
	a.style.SetTextAlign(css.Left)
}

//Right aligned text.
func (a Aligner) Right() {
	a.style.SetAlignSelf(css.FlexEnd)
	a.style.SetTextAlign(css.Right)
}

//Center aligned text.
func (a Aligner) Center() {
	a.style.SetAlignSelf(css.Center)
	a.style.SetTextAlign(css.Center)
}

//Align returns an aligner for aligning text.
func (text Seed) Align() Aligner {
	return Aligner{text.Style}
}
