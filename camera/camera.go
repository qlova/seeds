//Provides an embeded camera that can save snapshots.
package camera

import "strconv"

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Camera = seed.New()

	Camera.SetTag("video")
	Camera.SetAttributes("autoplay")

	Camera.OnReady(func(q script.Ctx) {
		var Camera = Camera.Ctx(q)
		q.Javascript(`
			if(navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
				navigator.mediaDevices.getUserMedia({ video: true }).then(function(stream) {
					` + Camera.Element() + `.srcObject = stream;
					` + Camera.Element() + `.play();
				});
			}
		
		`)
	})

	return Seed{Camera}
}

func AddTo(parent seed.Interface) Seed {
	var Camera = New()
	parent.Root().Add(Camera)
	return Camera
}

type Ctx struct {
	script.Seed
}

func (s Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{s.Seed.Ctx(q)}
}

type Image struct {
	Q script.Ctx

	script.Native
}

//Return a source that can be passed to image.SetSource
func (image Image) Source() script.String {
	return image.Q.Value(image.LanguageType().Raw() + ".toDataURL('image/jpg')").String()
}

func (camera Ctx) Capture() Image {
	var variable = script.Unique()

	camera.Q.Javascript(`let ` + variable + ";")
	camera.Q.Javascript(`{var canvas = document.createElement('canvas'); canvas.width = ` + camera.Element() + `.videoWidth; canvas.height = ` + camera.Element() + `.videoHeight;`)
	camera.Q.Javascript(`canvas.getContext('2d').drawImage(` + camera.Element() + `, 0, 0, canvas.width, canvas.height);`)
	camera.Q.Javascript(variable + ` = canvas}`)

	return Image{camera.Q, camera.Q.Value(variable).Native()}
}

func (image Image) AttachTo(request string, index int) string {
	return request + `.append("attachment-` + strconv.Itoa(index) + `-1", ` + image.LanguageType().Raw() + `.toBlob(function(b){}, 'image/jpeg', 0.95), 'camera.jpeg');`
}
