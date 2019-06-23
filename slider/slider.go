//Provides a basic slider that can be used as in input for a range of values.
package slider

import "strconv"
import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Slider = seed.New()

	Slider.SetTag("input")
	Slider.SetAttributes("type='range'")

	Slider.SetSize(seed.Auto, seed.Auto)

	Slider.Align(0)

	return Seed{Slider}
}

func AddTo(parent seed.Interface) Seed {
	var Slider = New()
	parent.Root().Add(Slider)
	return Slider
}

func (slider Seed) SetRequired() {
	slider.SetAttributes(slider.Attributes() + " required")
}

func (slider Seed) SetMax(max int) {
	slider.SetAttributes(slider.Attributes() + " max='" + strconv.Itoa(max) + "'")
}

type Script struct {
	script.Seed
}

func (slider Seed) Script(q script.Script) Script {
	return Script{slider.Seed.Script(q)}
}
