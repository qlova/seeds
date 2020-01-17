//Provide an image that can display various image types including png, jpeg and svg.
package image

import (
	"path"

	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/script"
)

//Seed is an image
type Seed struct {
	seed.Seed
}

//New returns a new image.
func New(source ...string) Seed {
	var Image = seed.New()

	Image.SetTag("img")
	if len(source) > 0 {
		var p = "/" + source[0]
		Image.Element.Set(html.Source, p)
		seed.NewAsset(p).AddTo(Image)
		Image.Set("alt", p[:len(p)-len(path.Ext(p))])
	}

	Image.Element.Set("crossorigin", "")

	return Seed{Image}
}

//AddTo parent.
func AddTo(parent seed.Interface, path ...string) Seed {
	var Image = New(path...)
	parent.Root().Add(Image)
	return Image
}

//Ctx is an image ctx.
type Ctx struct {
	script.Seed
}

//Ctx returns the image's ctx.
func (image Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{image.Seed.Ctx(q)}
}

const downloadJS = `
	async function downloadIMG(img, name) {
		
		function toDataURL(url) {
			return fetch(url).then((response) => {
					return response.blob();
				}).then(blob => {
					return URL.createObjectURL(blob);
				});
		}
		img.crossorigin = "anonymous";

		let link = document.createElement('a');
        link.download = name;
		link.href = await toDataURL(img.src);;
		
		document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
	}
`

//Download downloads the image and names it.
func (q Ctx) Download(name script.String) {
	q.Q.Require(downloadJS)
	q.Q.Javascript(`downloadIMG(%v, %v);`, q.Element(), name)
}
