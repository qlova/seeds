//Provide a header which is larger than standard text.
package header

import "github.com/qlova/seeds/text"

type Seed struct {
	text.Seed
}

func New(text ...string) Seed {
	var Header = text.New(text...)
	Header.SetTag("h1")
	return Seed{Header}
}

func AddTo(parent seed.Interface, text ...string) Widget {
	var Header = New(text...)
	parent.Root().Add(Header)
	return Header
}
