//Provides a swiper that can have multiple slides that can be swiped through.
package swiper

import "fmt"
import "encoding/json"

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

type Direction int

const Left Direction = -1
const Right Direction = 1

type Slide struct {
	index int

	seed.Seed
}

func init() {
	seed.Embed("/swiper.js", []byte(Javascript))
	seed.Embed("/swiper.css", []byte(CSS))
}

type Seed struct {
	slides int //The number of slides

	seed.Seed
	wrapper seed.Seed
}

func New(config ...Options) Seed {
	var Swiper = seed.New()

	Swiper.Require("swiper.js")
	Swiper.Require("swiper.css")

	wrapper := seed.AddTo(Swiper)
	wrapper.SetClass("swiper-wrapper")
	wrapper.SetRow()
	wrapper.SetSize(100, 100)

	pagination := seed.AddTo(Swiper)
	pagination.SetClass("swiper-pagination")

	var options string
	if len(config) > 0 {

		config[0].Pagination = PaginationOptions{
			Element: pagination,
		}

		var JSON, err = json.Marshal(config[0])
		if err == nil {
			options = string(JSON)
		}
	}

	Swiper.OnReady(func(q seed.Script) {
		q.Javascript(Swiper.Script(q).Element() + `.swiper = new Swiper('#` + Swiper.ID() + `', ` + options +
			`); window.addEventListener("resize", function() {
				setTimeout(function() {
					` + Swiper.Script(q).Element() + `.swiper.update();
				}, 250);
			}, false);window.addEventListener("orientationchange", function() {
				setTimeout(function() {
					` + Swiper.Script(q).Element() + `.swiper.update();
				}, 250);
			}, false);`)
	})

	return Seed{0, Swiper, wrapper}
}

func AddTo(parent seed.Interface, config ...Options) Seed {
	var Swiper = New(config...)
	parent.Root().Add(Swiper)
	return Swiper
}

func (swiper *Seed) NewSlide() Slide {
	var seed = seed.AddTo(swiper.wrapper)
	seed.SetClass("swiper-slide")

	seed.Set("display", "flex")
	seed.Set("align-items", "center")
	seed.Set("justify-items", "center")
	seed.Set("text-align", "center")
	seed.Set("flex-direction", "column")

	swiper.slides++

	return Slide{swiper.slides - 1, seed}
}

type Script struct {
	script.Seed
}

func (swiper Seed) Script(q script.Script) Script {
	return Script{swiper.Seed.Script(q)}
}

func (swiper Script) Update() {
	swiper.Q.Javascript(swiper.Element() + ".swiper.update(); window.dispatchEvent(new Event('resize'));")
}

func (swiper Script) Reset() {
	swiper.Q.Javascript(swiper.Element() + ".swiper.slideTo(0, 0);")
}

func (swiper Script) Goto(slide Slide) {
	swiper.Q.Javascript(swiper.Element() + ".swiper.slideTo(" + fmt.Sprint(slide.index) + ", 1000);")
}

func (swiper Script) Swipe(direction Direction) {
	if direction == Left {
		swiper.Q.Javascript(swiper.Element() + ".swiper.slidePrev();")
	}
	if direction == Right {
		swiper.Q.Javascript(swiper.Element() + ".swiper.slideNext();")
	}
}

//Are we on the last slide?
func (swiper Script) Left() script.Bool {
	return swiper.Q.Value(swiper.Element() + ".swiper.isBeginning").Bool()
}

//Are we on the first slide?
func (swiper Script) Right() script.Bool {
	return swiper.Q.Value(swiper.Element() + ".swiper.isEnd").Bool()
}
