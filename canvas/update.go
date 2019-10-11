package canvas

import (
	"github.com/qlova/seed/user"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script/js"
)

//Update allows remote control of the canvas.
type Update struct {
	seed.Update
}

//For returns an update to this seed.
func (seed Seed) For(u user.User) Update {
	return Update{seed.Seed.For(u)}
}

//Canvas2D resizes the canvas and returns the 2D ctx for doing graphical operations.
func (update Update) Canvas2D() Canvas2D {
	var value = js.ValueFromUpdate(update.Update)
	value.Set("width", value.Global().Get("innerWidth"))
	value.Set("height", value.Global().Get("innerHeight"))
	value = value.Call("getContext", value.String("2d")).Var("ctx")
	return Canvas2D{value}
}
