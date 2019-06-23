//Provides a button that can be clicked on to trigger an action.
package button

import "github.com/qlova/seed"

type Seed struct {
	seed.Seed
}

func New(label ...string) Seed {
	var Button = seed.New()

	Button.SetTag("button")

	Button.SetSize(seed.Auto, seed.Auto)

	if len(label) > 0 {
		Button.SetText(label[0])
	}

	return Seed{Button}
}

func AddTo(parent seed.Interface, label ...string) Seed {
	var Button = New(label...)
	parent.Root().Add(Button)
	return Button
}
