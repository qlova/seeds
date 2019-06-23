//Provides a canvas that can be rendered on with webGL.
package canvas

import qlova "github.com/qlova/script"
import "github.com/qlova/seed"
import "github.com/qlova/seed/gl"
import "github.com/qlova/seed/script"
import "github.com/qlova/seed/script/webgl"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Canvas = seed.New()

	Canvas.SetTag("canvas")
	Canvas.SetSize(seed.Auto, seed.Auto)

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

type Script struct {
	script.Seed
}

func (canvas Seed) Script(q seed.Script) Script {
	return Script{canvas.Seed.Script(q)}
}

func (canvas Script) OpenGL() webgl.Context {
	return webgl.NewContext(canvas.Seed)
}

func (canvas Script) Width() qlova.Float {
	return canvas.Q.Value(canvas.Element() + ".scrollWidth").Float()
}

func (canvas Script) Height() qlova.Float {
	return canvas.Q.Value(canvas.Element() + ".scrollHeight").Float()
}
