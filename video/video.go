package video

import (
	"strings"

	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

type Seed struct {
	seed.Seed
}

func New(path ...string) Seed {
	var Video = seed.New()

	Video.SetTag("video")
	if len(path) > 0 {
		if strings.HasPrefix(path[0], "https://") {
			var url = path[0]
			if strings.HasPrefix(url, `https://player.vimeo.com`) {
				Video.SetTag("iframe")
				Video.Element.Set(html.Source, url)
				Video.Element.Set(html.AllowFullscreen)
				Video.Element.Set(html.Frameborder, "0")
				return Seed{Video}
			}
		}
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
