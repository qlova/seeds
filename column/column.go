//Provides a column that stacks children vertically.
package column

import "github.com/qlova/seed"
import "github.com/qlova/seed/style/css"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Column = seed.New()

	Column.SetDisplay(css.Flex)
	Column.SetFlexDirection(css.Column)

	return Seed{Column}
}

func AddTo(parent seed.Interface) Seed {
	var Column = New()
	parent.Root().Add(Column)
	return Column
}
