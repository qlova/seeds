//Provides a checkbox that can in either a checked or unchecked state.
package checkbox

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Checkbox = seed.New()

	Checkbox.SetTag("input")
	Checkbox.SetAttributes("type='checkbox'")

	return Seed{Checkbox}
}

func AddTo(parent seed.Interface) Seed {
	var Checkbox = New()
	parent.Root().Add(Checkbox)
	return Checkbox
}

type Script struct {
	script.Seed
}

func (checkbox Seed) Script(q seed.Script) Script {
	return Script{checkbox.Seed.Script(q)}
}

func (checkbox Script) Checked() script.Bool {
	return checkbox.Q.Value(checkbox.Element()+".checked").Bool()
}

func (checkbox Script) Check() script.Bool {
	checkbox.Q.Javascript(checkbox.Element()+".check = true;")
}

func (checkbox Script) Uncheck() script.Bool {
	checkbox.Q.Javascript(checkbox.Element()+".check = false;")
}
