//Provides emedded documents (such as PDF)
package document

import "github.com/qlova/seed"

type Seed struct {
	seed.Seed
}

func New(path ...string) Seed {
	var Document = seed.New()

	Document.SetTag("embed")
	if len(path) > 0 {
		Document.SetAttributes("src='" + path[0] + "'")
		seed.NewAsset(path[0]).AddTo(Document)
	}

	return Seed{Document}
}

func AddTo(parent seed.Interface, path ...string) Seed {
	var Document = New(path...)
	parent.Root().Add(Document)
	return Document
}
