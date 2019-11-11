//Package base provides an html base element.
package base

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

//Seed is a base.
type Seed struct {
	seed.Seed
}

//New returns a new base.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("base")
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}

/*
SetHypertextReference sets the base URL to be used
throughout the document for relative URLs.
Absolute and relative URLs are allowed.
*/
func (base Seed) SetHypertextReference(url string) {
	base.Element.Set(html.HypertextReference, url)
}

/*
SetTarget sets where to display the linked URL,
as the name for a browsing context (a tab, window, or <iframe>).
*/
func (base Seed) SetTarget(target html.TargetType) {
	base.Element.Set(html.Target, string(target))
}
