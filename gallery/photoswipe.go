//Provide a photoswipe gallery that can lazy load images from the server.
package gallery

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func init() {
	seed.Embed("/photoswipe.js", []byte(Javascript))
	seed.Embed("/photoswipe.css", []byte(CSS))
	seed.Embed("/photoswipe-ui.js", []byte(UI))
}

type Seed struct {
	seed.Seed
}

func getImageDimension(imagePath string) string {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}

	file.Close()

	return fmt.Sprint("w:", image.Width, ",h:", image.Height)
}

//Returns gallery that displays 'local' images (in the assets directory).
func New(images ...string) Seed {
	var Gallery = seed.New()

	Gallery.Require("photoswipe.js")
	Gallery.Require("photoswipe.css")
	Gallery.Require("photoswipe-ui.js")

	Gallery.OnReady(func(q script.Ctx) {
		q.Javascript(Gallery.Ctx(q).Element() + ".items = [")
		for i, img := range images {

			var dimensions = getImageDimension(seed.Dir + "/assets/" + img)

			q.Javascript(`{src:"` + img + `", ` + dimensions + ` }`)
			if i < len(images)-1 {
				q.Javascript(`,`)
			}
		}
		q.Javascript("];")
	})

	return Seed{Gallery}
}

func AddTo(parent seed.Interface, images ...string) Seed {
	var Gallery = New(images...)
	parent.Root().Add(Gallery)
	return Gallery
}

type Ctx struct {
	script.Seed
}

func (gallery Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{gallery.Seed.Ctx(q)}
}

func (gallery Ctx) Open() {
	gallery.Q.Javascript(`if (document.querySelectorAll(".pswp").length == 0) {
		document.body.insertAdjacentHTML("beforeend", '<div class="pswp" tabindex="-1" role="dialog" aria-hidden="true"><!-- Background of PhotoSwipe. Its a separate element as animating opacity is faster than rgba(). --> <div class="pswp__bg"></div><div class="pswp__scroll-wrap"><!-- Container that holds slides. PhotoSwipe keeps only 3 of them in the DOM to save memory. Dont modify these 3 pswp__item elements, data is added later on. --> <div class="pswp__container"> <div class="pswp__item"></div><div class="pswp__item"></div><div class="pswp__item"></div></div><div class="pswp__ui pswp__ui--hidden"> <div class="pswp__top-bar"> <div class="pswp__counter"></div><button class="pswp__button pswp__button--close" title="Close (Esc)"></button><button class="pswp__button pswp__button--fs" title="Toggle fullscreen"></button> <button class="pswp__button pswp__button--zoom" title="Zoom in/out"></button> <div class="pswp__preloader"> <div class="pswp__preloader__icn"> <div class="pswp__preloader__cut"> <div class="pswp__preloader__donut"></div></div></div></div></div><div class="pswp__caption"> <div class="pswp__caption__center"></div></div></div></div></div>');
	}`)
	gallery.Q.Javascript(`ActivePhotoSwipe = new PhotoSwipe(document.querySelectorAll(".pswp")[0], PhotoSwipeUI_Default, ` + gallery.Element() + ".items, {history:false}); ActivePhotoSwipe.init(); ActivePhotoSwipe.listen('close', function() { ActivePhotoSwipe = null; });")
}
