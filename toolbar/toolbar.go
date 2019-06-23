//Provides basic toolbar that is sticky.
package toolbar

import "github.com/qlova/seed"
import "github.com/qlova/seed/style/css"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var ToolBar = seed.New()

	ToolBar.Stylable.Set("display", "flex")
	ToolBar.Stylable.Set("position", "fixed")
	ToolBar.SetFlexDirection(css.Row)
	ToolBar.SetBottom(css.Zero)
	ToolBar.SetLeft(css.Zero)

	ToolBar.SetWidth(100)
	ToolBar.SetHeight(2 * seed.Em)

	ToolBar.Set("justify-content", "space-evenly")

	return Seed{ToolBar}
}

func AddTo(parent seed.Interface) Seed {
	var ToolBar = New()
	ToolBar.AddTo(parent)
	return ToolBar
}
