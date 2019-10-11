//Provides a canvas that can be rendered on with webGL.
package canvas

import (
	qlova "github.com/qlova/script"
	"github.com/qlova/seed"
	"github.com/qlova/seed/gl"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/script/webgl"
	"github.com/qlova/seed/unit"
)

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Canvas = seed.New()

	Canvas.SetTag("canvas")
	Canvas.SetSize(unit.Auto, unit.Auto)

	return Seed{Canvas}
}

func AddTo(parent seed.Interface) Seed {
	var Canvas = New()
	parent.Root().Add(Canvas)
	return Canvas
}

//Return an OpenGL context for this canvas.
func (s Seed) OpenGL() gl.Context {
	return gl.NewContext(s.Seed)
}

func (canvas Seed) OnDraw(f func(q script.Ctx)) {
	canvas.On("draw", func(q script.Ctx) {
		f(q)
		q.Javascript(`requestAnimationFrame(` + canvas.Ctx(q).Element() + `.ondraw);`)
	})
	canvas.OnReady(func(q script.Ctx) {
		q.Javascript(`requestAnimationFrame(` + canvas.Ctx(q).Element() + `.ondraw);`)
	})
}

type Ctx struct {
	script.Seed
}

func (canvas Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{canvas.Seed.Ctx(q)}
}

func (canvas Ctx) OpenGL() webgl.Context {
	return webgl.NewContext(canvas.Seed)
}

func (canvas Ctx) Width() qlova.Float {
	return canvas.Q.Value(canvas.Element() + ".scrollWidth").Float()
}

func (canvas Ctx) Height() qlova.Float {
	return canvas.Q.Value(canvas.Element() + ".scrollHeight").Float()
}
