//Provides an expander that will use up empty space at a specified ratio.
package expander

import "github.com/qlova/seed"

type Seed struct {
	seed.Seed
}

func New(ratio ...float64) Seed {
	var Expander = seed.New()

	if len(ratio) > 0 {
		Expander.SetExpand(ratio[0])
	} else {
		Expander.SetExpand(1)
	}

	return Seed{Expander}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, ratio ...float64) Seed {
	var Expander = New(ratio...)
	parent.Root().Add(Expander)
	return Expander
}
