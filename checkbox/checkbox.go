//Provides a checkbox that can in either a checked or unchecked state.
package checkbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style/css"
)

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Checkbox = seed.New()

	Checkbox.SetTag("input")
	Checkbox.Element.Set(`type`, `checkbox`)

	return Seed{Checkbox}
}

func AddTo(parent seed.Interface) Seed {
	var Checkbox = New()
	parent.Root().Add(Checkbox)
	return Checkbox
}

//SetReadOnly sets this seed to be read only.
func (checkbox Seed) SetReadOnly() {
	checkbox.CSS().SetPointerEvents(css.None)
	checkbox.Seed.SetAttributes(checkbox.Seed.Attributes() + " readonly")
}

type Ctx struct {
	script.Seed
}

func (checkbox Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{checkbox.Seed.Ctx(q)}
}

func (checkbox Ctx) Checked() script.Bool {
	return checkbox.Q.Value(checkbox.Element() + ".checked").Bool()
}

func (checkbox Ctx) Check() {
	checkbox.Q.Javascript(checkbox.Element() + ".checked = true;")
}

func (checkbox Ctx) Toggle() {
	checkbox.Q.Javascript(checkbox.Element() + ".checked = !" + checkbox.Element() + ".checked ;")
}

func (checkbox Ctx) Uncheck() {
	checkbox.Q.Javascript(checkbox.Element() + ".checked = false;")
}
