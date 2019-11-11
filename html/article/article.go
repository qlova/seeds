//Package article provides a html article element.
package article

import (
	"github.com/qlova/seed"
)

//Seed is a article.
type Seed struct {
	seed.Seed
}

//New returns a new article.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("article")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}