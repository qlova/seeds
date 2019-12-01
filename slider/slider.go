//Provides a basic slider that can be used as in input for a range of values.
package slider

import (
	"image/color"
	"strconv"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/unit"
)

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Slider = seed.New()

	Slider.SetTag("input")
	Slider.SetAttributes("type='range'")

	Slider.SetSize(unit.Auto, unit.Auto)

	Slider.Align().Center()

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

//Thumb returns the style of the thumb, this can be modified.
func (slider Seed) Thumb() style.Style {
	slider.CSS().Set("-webkit-appearance", "none")
	var webkit = slider.Select("::-webkit-slider-thumb")
	webkit.CSS().Set("-webkit-appearance", "none")
	var mozilla = slider.Select("::-moz-range-thumb")
	return style.Compose(webkit, mozilla)
}

//SetColors sets the color of the two side of this range.
func (slider Seed) SetColors(a, b color.Color) {
	slider.SetColor(b)
	slider.Select("::-moz-range-progress").SetColor(a)

	slider.OnReady(func(q script.Ctx) {
		q.Javascript(`if ('webkitRequestAnimationFrame' in window) {
			let element = %v;
			let old = element.oninput;
			element.oninput = function() {
				let value = 100*(element.value/element.max);

				element.style.background = 'linear-gradient(to right, %[2]v 0%%, %[2]v '+value +'%%, %[3]v ' + value + '%%, %[3]v 100%%)';

				if (old) old();
			};
			element.oninput();
		}`, slider.Ctx(q).Element(), string(css.Colour(a)), string(css.Colour(b)))
	})
}

type Ctx struct {
	script.Seed
}

func (slider Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{slider.Seed.Ctx(q)}
}
