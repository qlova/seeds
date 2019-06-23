//Provide a seperator that renders as a line that can visually break different sections of a container.
package line

import "image/color"
import "github.com/qlova/seed"
import "github.com/qlova/seed/style/css"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Line = seed.New()

	Line.SetTag("hr")
	Line.SetSize(seed.Auto, seed.Auto)
	Line.Set("border-style", "solid")

	return Seed{Line}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface) Seed {
	var Line = New()
	parent.Root().Add(Line)
	return Line
}

func (line Seed) SetColor(c color.Color) {
	line.SetBorderColor(css.Colour(c))
}
