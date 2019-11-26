//Provide an mesh that can display various mesh types including gltf, glb and obj.
package mesh

import (
	"strconv"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

func init() {
	seed.Embed("/three.js", []byte(ThreeJS))
	seed.Embed("/three.js_gltf.js", []byte(GLTFLoader))
	seed.Embed("/three.js_orbit.js", []byte(OrbitJS))
}

//Seed is a mesh.
type Seed struct {
	seed.Seed
}

//New returns a new mesh.
func New(source ...string) Seed {
	var Mesh = seed.New()
	Mesh.Require("/three.js")
	Mesh.Require("/three.js_gltf.js")
	Mesh.Require("/three.js_orbit.js")

	Mesh.SetTag("canvas")

	if len(source) > 0 {
		var filename = source[0]
		seed.NewAsset(filename).AddTo(Mesh)

		Mesh.OnReady(func(q script.Ctx) {
			q.Javascript(`
				var scene = new THREE.Scene();
				var camera = new THREE.PerspectiveCamera( 70, 2, 1, 1000 );
			
				let canvas = %v;
				var renderer = new THREE.WebGLRenderer( { alpha: true, canvas:canvas } );
				renderer.setSize(canvas.innerWidth, canvas.innerHeight);
				renderer.setPixelRatio( window.devicePixelRatio );

				var controls = new THREE.OrbitControls( camera, renderer.domElement );
				camera.position.set( 0, 0, 20 );
				controls.update();
				controls.minDistance = 20;
				controls.enablePan = false;

				function resizeCanvasToDisplaySize() {
					const canvas = renderer.domElement;
					// look up the size the canvas is being displayed
					const width = canvas.clientWidth;
					const height = canvas.clientHeight;
				  
					// adjust displayBuffer size to match
					if (canvas.width !== width || canvas.height !== height) {
					  // you must pass false here or three.js sadly fights the browser
					  renderer.setSize(width, height, false);
					  camera.aspect = width / height;
					  camera.updateProjectionMatrix();
				  
					  // update any render target sizes here
					}
				  }

				scene.add(new THREE.AmbientLight( 0x404040 ));

				// Instantiate a loader
				var loader = new THREE.GLTFLoader();

				loader.load(%v,
					function ( gltf ) {
				
						scene.add( gltf.scene );

						gltf.scene.position.y -= 5;
				
						function animate(time) {
							resizeCanvasToDisplaySize();
						  
							controls.update();
							renderer.render(scene, camera);
							requestAnimationFrame(animate);
						  }
						  
						  requestAnimationFrame(animate);
				
					},
					// called while loading is progressing
					function ( xhr ) {
				
						console.log( ( xhr.loaded / xhr.total * 100 ) + '% loaded' );
				
					},
					// called when loading has errors
					function ( error ) {
				
						console.log( 'An error happened', error );
				
					}
				);

				
			`, Mesh.Ctx(q).Element(), strconv.Quote(filename))

		})
	}

	return Seed{Mesh}
}

//AddTo parent
func AddTo(parent seed.Interface, path ...string) Seed {
	var Mesh = New(path...)
	parent.Root().Add(Mesh)
	return Mesh
}
