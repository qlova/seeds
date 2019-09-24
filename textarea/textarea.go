//Provides basic textarea for users to edit multiline text.
package textarea

import (
	"github.com/qlova/seed"
)

type Seed struct {
	seed.Seed
}

func New() Seed {
	var TextArea = seed.New()

	TextArea.SetTag("textarea")
	TextArea.SetAttributes("data-gramm_editor=false")
	TextArea.CSS().Set("resize", "none")

	return Seed{TextArea}
}

func AddTo(parent seed.Interface) Seed {
	var TextArea = New()
	parent.Root().Add(TextArea)
	return TextArea
}

func (textarea Seed) SetRequired() {
	textarea.SetAttributes(textarea.Attributes() + " required")
}

func (textarea Seed) SetPlaceholder(placeholder string) {
	textarea.SetAttributes(textarea.Attributes() + " placeholder='" + placeholder + "' ")
}
