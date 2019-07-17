//Provide a header which is larger than standard text.
package header

import "github.com/qlova/seed"
import "github.com/qlova/seeds/text"

type Seed struct {
	text.Seed
}

func New(s ...string) Seed {
	var Header = text.New(s...)
	Header.SetTag("h1")
	return Seed{Header}
}

func AddTo(parent seed.Interface, s ...string) Seed {
	var Header = New(s...)
	parent.Root().Add(Header)
	return Header
}
