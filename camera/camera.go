//Provides an embeded camera that can save snapshots.
package camera

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

type Seed struct {
	seed.Seed
}

func New() Seed {
	var Camera = seed.New()

	Camera.SetTag("video")
	Camera.SetAttributes("autoplay")

	Camera.OnReady(func(q seed.Script) {
		var Camera = Camera.Script(q)
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

type Script struct {
	script.Seed
}

func (s Seed) Script(q seed.Script) Script {
	return Script{s.Seed.Script(q)}
}

type Image struct {
	Q script.Script

	script.Native
}

//Return a source that can be passed to image.SetSource
func (image Image) Source() script.String {
	return image.Q.Value(image.LanguageType().Raw()+".toDataURL('image/jpg')").String()
}

func (camera Script) Capture() Image {
	var variable = script.Unique()

	camera.Q.Javascript(`let `+variable+";")
	camera.Q.Javascript(`{var canvas = document.createElement('canvas'); canvas.width = `+camera.Element()+`.videoWidth; canvas.height = `+camera.Element()+`.videoHeight;`)
	camera.Q.Javascript(`canvas.getContext('2d').drawImage(`+camera.Element()+`, 0, 0, canvas.width, canvas.height);`)
	camera.Q.Javascript(variable+` = canvas}`)

	return Image{camera.Q, camera.Q.Value(variable).Native()}
}
