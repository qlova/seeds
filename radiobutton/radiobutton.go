//Provides a radiobutton that can in either a checked or unchecked state.
package radiobutton

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style/css"
)

type Seed struct {
	seed.Seed
}

func New(name string) Seed {
	var RadioButton = seed.New()

	RadioButton.SetTag("input")
	RadioButton.Element.Set(`type`, `radio`)
	RadioButton.Element.Set(`name`, name)

	return Seed{RadioButton}
}

func AddTo(parent seed.Interface, name string) Seed {
	var RadioButton = New(name)
	parent.Root().Add(RadioButton)
	return RadioButton
}

//SetReadOnly sets this seed to be read only.
func (RadioButton Seed) SetReadOnly() {
	RadioButton.CSS().SetPointerEvents(css.None)
	RadioButton.Seed.SetAttributes(RadioButton.Seed.Attributes() + " readonly")
}

type Ctx struct {
	script.Seed
}

func (RadioButton Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{RadioButton.Seed.Ctx(q)}
}

func (RadioButton Ctx) Checked() script.Bool {
	return RadioButton.Q.Value(RadioButton.Element() + ".checked").Bool()
}

func (RadioButton Ctx) Check() {
	RadioButton.Q.Javascript(RadioButton.Element() + ".checked = true;")
}

func (RadioButton Ctx) Toggle() {
	RadioButton.Q.Javascript(RadioButton.Element() + ".checked = !" + RadioButton.Element() + ".checked ;")
}

func (RadioButton Ctx) Uncheck() {
	RadioButton.Q.Javascript(RadioButton.Element() + ".checked = false;")
}
