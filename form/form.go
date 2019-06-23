//Provides a form that can be used to group together input elements in order to report their validity.
package form

import "github.com/qlova/seed"
import "github.com/qlova/seed/style/css"
import "github.com/qlova/seed/script"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Form = seed.New()

	Form.SetTag("form")
	Form.SetAttributes(`onsubmit="return false;"`)

	Form.SetDisplay(css.Flex)
	Form.SetFlexDirection(css.Column)

	return Seed{Form}
}

func AddTo(parent seed.Interface) Seed {
	var Form = New()
	parent.Root().Add(Form)
	return Form
}

type Script struct {
	script.Seed
}

func (form Seed) Script(q script.Script) Script {
	return Script{form.Seed.Script(q)}
}

func (form Script) Invalid() script.Bool {
	return form.Q.Value("!" + form.Element() + ".reportValidity()").Bool()
}
