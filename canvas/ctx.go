package canvas

import (
	"image/color"

	"github.com/qlova/seed/script/js"

	"github.com/qlova/seed/style/css"
)

//Canvas2D is an interface to graphical output for this canvas.
type Canvas2D struct {
	js.Value
}

//SetFillColor sets the color of subsequent fills.
func (ctx Canvas2D) SetFillColor(color color.Color) {
	ctx.Value.Set("fillStyle", ctx.String(css.Colour(color).String()))
}

//FillRect fills a rectangle of the specified bounds.
func (ctx Canvas2D) FillRect(x, y, w, h float64) {
	ctx.Value.Run("fillRect", ctx.Number(x), ctx.Number(y), ctx.Number(w), ctx.Number(h))
}
