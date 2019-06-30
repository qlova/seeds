//Provides a basic circle renderer, can contain children.
package circle

import "github.com/qlova/seed"

type Seed struct {
	seed.Seed
}

func New(size ...complex128) Seed {
	var Circle = seed.New()

	if len(size) > 0 {
		Circle.SetSize(size[0], size[0])
	}
	
	Circle.SetRoundedCorners(50)

	return Seed{Circle}
}

func AddTo(parent seed.Interface, size ...complex128) Seed {
	var Circle = New(size...)
	parent.Root().Add(Circle)
	return Circle
}
