package video

import "github.com/qlova/seed"

type Seed struct {
	seed.Seed
}

func New(path ...string) Seed {
	var Video = seed.New()

	Video.SetTag("video")
	if len(path) > 0 {
		Video.SetAttributes("src='" + path[0] + "' playsinline preload='auto'")
		seed.NewAsset(path[0]).AddTo(Video)
	} else {
		Video.SetAttributes("playsinline preload='auto'")
	}

	return Seed{Video}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, path ...string) Seed {
	var Video = New(path...)
	parent.Root().Add(Video)
	return Video
}
