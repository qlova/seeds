//Provides a linke that can be clicked on to launch a website.
package link

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

type Seed struct {
	seed.Seed
}

func New(url ...string) Seed {
	var Link = seed.New()
	Link.SetTag("a")

	if len(url) > 0 {
		Link.SetAttributes("href='" + url[0] + "'")
	} else {
		Link.SetAttributes("href='#'")
	}

	return Seed{Link}
}

func AddTo(parent seed.Interface, url ...string) Seed {
	var Link = New(url...)
	parent.Root().Add(Link)
	return Link
}

type Ctx struct {
	script.Seed
}

func (link Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{link.Seed.Ctx(q)}
}

func (link Ctx) Target() script.String {
	return link.Q.Value(link.Element() + `.href`).String()
}

func (link Ctx) SetTarget(target script.String) {
	link.Q.Javascript(link.Element() + `.href = ` + string(link.Q.Raw(target)) + ";")
}

func (link Ctx) Open() {
	link.Q.Website(link.Target()).Open()
}
