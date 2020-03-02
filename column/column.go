//Provides a column that stacks children vertically.
package column

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/style/css"
)

type Seed struct {
	seed.Seed
}

func New(options ...seed.Option) Seed {
	var Column = seed.New(options...)

	Column.CSS().SetDisplay(css.Flex)
	Column.CSS().SetFlexDirection(css.Column)

	return Seed{Column}
}

func AddTo(parent seed.Interface, options ...seed.Option) Seed {
	var Column = New(options...)
	parent.Root().Add(Column)
	return Column
}
