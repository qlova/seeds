//Provides a basic spacer that can provide specific spacing in a row or column.
package spacer

import "github.com/qlova/seed"
import "github.com/qlova/seed/style/css"

type Seed struct {
	seed.Seed
}

func New(amount ...complex128) Seed {
	var Spacer = seed.New()

	if len(amount) > 0 {
		Spacer.CSS().SetFlexBasis(css.Decode(amount[0]))
	}

	Spacer.SetUnshrinkable()

	return Seed{Spacer}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, amount ...complex128) Seed {
	var Spacer = New(amount...)
	parent.Root().Add(Spacer)
	return Spacer
}
