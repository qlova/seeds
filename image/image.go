//Provide an image that can display various image types including png, jpeg and svg.
package image

import (
	"path"

	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

type Seed struct {
	seed.Seed
}

func New(source ...string) Seed {
	var Image = seed.New()

	Image.SetTag("img")
	if len(source) > 0 {
		var p = source[0]
		Image.Element.Set(html.Source, p)
		seed.NewAsset(p).AddTo(Image)
		Image.Set("alt", p[:len(p)-len(path.Ext(p))])
	}

	return Seed{Image}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, path ...string) Seed {
	var Image = New(path...)
	parent.Root().Add(Image)
	return Image
}
