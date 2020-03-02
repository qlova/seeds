//Provides a textbox that only accepts email addresses.
package emailbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seeds/textbox"
)

type Seed struct {
	textbox.Seed
}

func New() Seed {
	var EmailBox = textbox.New()
	EmailBox.SetAttributes(`type="email"`)

	return Seed{EmailBox}
}

func AddTo(parent seed.Interface) Seed {
	var EmailBox = New()
	parent.Root().Add(EmailBox)
	return EmailBox
}

func (emailbox Seed) ValueFromCtx(ctx script.AnyCtx) script.Value {
	q := script.CtxFromAnyCtx(ctx)
	return emailbox.Ctx(q).Value()
}
