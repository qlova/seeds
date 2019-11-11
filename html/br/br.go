//Package br provides a html br element.
package br

import (
	"github.com/qlova/seed"
)

//Seed is a br.
type Seed struct {
	seed.Seed
}

//New returns a new br.
func New() Seed {
	var Text = seed.New()
	Text.SetTag("br")			
	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}