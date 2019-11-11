//Package table provides a html table element.
package table

import (
	"github.com/qlova/seed"
)

//Seed is a table.
type Seed struct {
	seed.Seed
}

//New returns a new table.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("table")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}