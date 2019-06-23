//Provide a listbox where a user can select from multiple items.
package listbox

import "fmt"
import "github.com/qlova/seed"

type Seed struct {
	seed.Seed
}

func New(values ...string) Seed {
	var ListBox = seed.New()
	ListBox.SetTag("select")

	var content string

	for _, value := range values {
		content += fmt.Sprint("<option value='", value, "'>", value, "</option>")
	}

	ListBox.SetContent(content)

	return Seed{ListBox}
}

func AddTo(parent seed.Interface, values ...string) Seed {
	var listbox = New(values...)
	parent.Root().Add(listbox)
	return listbox
}
