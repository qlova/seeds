//Provide a seperator that renders as a line that can visually break different sections of a container.
package line

import "github.com/qlova/seed"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Line = seed.New()

	Line.SetTag("hr")
	Line.CSS().Set("border", "none")
	Line.SetSize(100, 1*seed.Px)

	return Seed{Line}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface) Seed {
	var Line = New()
	parent.Root().Add(Line)
	return Line
}
