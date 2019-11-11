//Package video provides a html video element.
package video

import (
	"github.com/qlova/seed"
)

//Seed is a video.
type Seed struct {
	seed.Seed
}

//New returns a new video.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("video")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}