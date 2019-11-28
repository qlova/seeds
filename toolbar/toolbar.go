//Provides basic toolbar that is sticky.
package toolbar

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/unit"
)

type Seed struct {
	seed.Seed
}

func New() Seed {
	var ToolBar = seed.New()

	ToolBar.CSS().Set("display", "flex")
	ToolBar.CSS().Set("position", "fixed")
	ToolBar.CSS().SetFlexDirection(css.Row)
	ToolBar.CSS().SetBottom(css.Zero)
	ToolBar.CSS().SetLeft(css.Zero)

	ToolBar.SetWidth(100)
	ToolBar.SetHeight(2 * unit.Em)

	ToolBar.Set("justify-content", "space-evenly")

	return Seed{ToolBar}
}

func AddTo(parent seed.Interface) Seed {
	var ToolBar = New()
	ToolBar.AddTo(parent)
	return ToolBar
}
