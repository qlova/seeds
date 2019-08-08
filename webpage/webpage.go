package webpage

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

type Seed struct {
	seed.Seed
}

func New(path ...string) Seed {
	var WebPage = seed.New()

	WebPage.SetTag("iframe")
	if len(path) > 0 {
		WebPage.Element.Set(html.Source, path[0])
		WebPage.Element.Set(html.Frameborder, "0")
	}

	return Seed{WebPage}
}

//Create a new WebPage widget and add it to the provided parent.
func AddTo(parent seed.Interface, path ...string) Seed {
	var WebPage = New(path...)
	parent.Root().Add(WebPage)
	return WebPage
}
