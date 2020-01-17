//Package anime performs automatic animations on elements.
package anime

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

func init() {
	seed.Embed("/flipping.js", []byte(js))
}

//Push should be called before a layout change.
func Push(q script.Ctx) {
	q.Javascript(`flipping.read();`)
}

//Pop should be called after a layout change.
func Pop(q script.Ctx) {
	q.Javascript(`try { flipping.flip(); } catch(error) {}`)
}

//Do creates animates any layout changes made during the call to 'f'
func Do(q script.Ctx, f func()) {
	q.Javascript(`flipping.read();`)
	f()
	q.Javascript(`try { flipping.flip(); } catch(error) {}`)
}

//Animate an element.
func Animate(element seed.Interface) {
	var root = element.Root()
	root.Require("/flipping.js")
	root.Element.Set("data-flip-key", root.ID())
}

//Track an element between multiple pages.
func Track(element seed.Interface, trackingCode string) {
	var root = element.Root()
	root.Require("/flipping.js")
	root.Element.Set("data-flip-key", trackingCode)
}
