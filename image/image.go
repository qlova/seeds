//Provide an image that can display various image types including png, jpeg and svg.
package image

import "github.com/qlova/seed"

type Seed struct {
	seed.Seed
}

func New(path ...string) Seed {
	var Image = seed.New()

	Image.SetTag("img")
	if len(path) > 0 {
		Image.SetAttributes("src='" + path[0] + "'")
		seed.NewAsset(path[0]).AddTo(Image)
	}

	return Seed{Image}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, path ...string) Seed {
	var Image = New(path...)
	parent.Root().Add(Image)
	return Image
}
