//Provides a simple text editor with line numbers.
package editor

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

func init() {
	seed.Embed("/ace.js", []byte(Javascript))
}

type Seed struct {
	seed.Seed
}

//Returns a full-featured text editor with line numbers and optional syntax highlighting.
func New(text ...string) Seed {
	var Editor = seed.New()
	Editor.SetTag("pre")

	if len(text) > 0 {
		Editor.SetText(text[0])
	}

	Editor.Require("ace.js")

	Editor.OnReady(func(q script.Ctx) {
		q.Javascript(`let editor = ace.edit("` + Editor.ID() + `"); editor.setShowPrintMargin(false); document.getElementById("` + Editor.ID() + `").editor = editor;`)
	})

	return Seed{Editor}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, path ...string) Seed {
	var Editor = New(path...)
	parent.Root().Add(Editor)
	return Editor
}

type Editor struct {
	script.Seed
}

func (editor Editor) Open(f script.File) {
	editor.Javascript(`var reader = new FileReader(); reader.onload = function(e) { var data = e.target.result; get("` + editor.ID + `").editor.setValue(data); }; reader.readAsText(` + editor.Q.Raw(f) + `);`)
}
