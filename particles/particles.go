package particles

import (
	"encoding/json"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

func init() {
	seed.Embed("/particles.js", []byte(JS))
}

type Seed struct {
	seed.Seed
}

func New(options Options) Seed {
	var container = seed.New()

	container.Require("particles.js")

	var JSON, err = json.Marshal(options)
	if err != nil {
		panic("Invalid Particle options")
	}

	container.OnReady(func(q script.Ctx) {
		q.Javascript("particlesJS('" + container.ID() + "', " + string(JSON) + ")")
	})

	return Seed{container}
}

func AddTo(parent seed.Interface, options Options) Seed {
	var Particles = New(options)
	parent.Root().Add(Particles)
	return Particles
}
