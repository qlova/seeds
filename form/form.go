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

	Form.CSS().SetDisplay(css.Flex)
	Form.CSS().SetFlexDirection(css.Column)

	return Seed{Form}
}

func AddTo(parent seed.Interface) Seed {
	var Form = New()
	parent.Root().Add(Form)
	return Form
}

type Ctx struct {
	script.Seed
}

func (form Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{form.Seed.Ctx(q)}
}

func (form Ctx) Invalid() script.Bool {
	return form.Q.Value("!" + form.Element() + ".reportValidity()").Bool()
}
