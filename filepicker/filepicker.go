//Provides a filepicker that allows the user to select a file from their filesystem.
package filepicker

import (
	"strconv"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

type Seed struct {
	seed.Seed
}

func New(types ...string) Seed {
	var FilePicker = seed.New()
	FilePicker.SetTag("Input")

	if len(types) > 0 {
		FilePicker.SetAttributes(`type="file" accept="` + types[0] + `"`)
	} else {
		FilePicker.SetAttributes(`type="file" accept="*"`)
	}

	return Seed{FilePicker}
}

func AddTo(parent seed.Interface, types ...string) Seed {
	var FilePicker = New(types...)
	parent.Root().Add(FilePicker)
	return FilePicker
}

type Ctx struct {
	script.Seed
}

func (filepicker Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{filepicker.Seed.Ctx(q)}
}

func (filepicker Ctx) AttachTo(request string, index int) string {

	return "for (let i = 0; i < " + filepicker.Element() + ".files.length; i++) " + request +
		`.append("attachment-` + strconv.Itoa(index) + `-"+(i+1), ` + filepicker.Element() + `.files[i], ` + filepicker.Element() + `.files[i].name);`
}

//File returns the first file.
func (filepicker Ctx) File() script.File {
	return filepicker.Q.Value(filepicker.Element() + ".files[0]").File()
}

//Clear clears the file.
func (filepicker Ctx) Clear() script.File {
	return filepicker.Q.Value(filepicker.Element() + ".value = '';").File()
}
