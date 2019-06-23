//Provides a row that stacks children horizontally.
package row

import "github.com/qlova/seed"
import "github.com/qlova/seed/style/css"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Row = seed.New()

	Row.SetDisplay(css.Flex)
	Row.SetFlexDirection(css.Row)

	return Seed{Row}
}

func AddTo(parent seed.Interface) Seed {
	var Row = New()
	parent.Root().Add(Row)
	return Row
}
