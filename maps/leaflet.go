//Provide an embedded leaflet map.
package maps

import (
	"encoding/json"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

func init() {
	seed.Embed("/leaflet.js", []byte(Javascript))
	seed.Embed("/leaflet.css", []byte(Stylesheet))
}

type Seed struct {
	seed.Seed
}

//Returns a full-featured map.
func New(config ...Options) Seed {
	var Maps = seed.New()

	var options string
	if len(config) > 0 {
		var JSON, err = json.Marshal(config[0])
		if err == nil {
			options = string(JSON)
		}
	}

	Maps.Require("leaflet.js")
	Maps.Require("leaflet.css")

	Maps.SetSize(100, 100)

	Maps.OnReady(func(q script.Ctx) {
		q.Javascript(`let map = L.map("` + Maps.ID() + `", ` + options + `); L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', { maxZoom: 19, attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>' }).addTo(map); get("` + Maps.ID() + `").map = map;`)
	})

	return Seed{Maps}
}

func AddTo(parent seed.Interface, config ...Options) Seed {
	var Maps = New(config...)
	parent.Root().Add(Maps)
	return Maps
}

type Ctx struct {
	script.Seed
}

func (maps Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{maps.Seed.Ctx(q)}
}

func (maps Ctx) FlyTo(location script.Location) {
	var raw = maps.Q.Raw(location)
	maps.Q.Javascript(maps.Element() + ".map.flyTo(L.latLng(" + raw + ".coords.latitude, " + raw + ".coords.longitude))")
}

func (maps Ctx) Location() script.Location {
	var q = maps.Q
	q.Require(`function maps_get_location_of(map) {
		let center = map.getCenter();
		return {
			coord: {
				latitude: center.lat,
				longitude: center.lng,
			},
		}
	}`)
	return q.Value(`maps_get_location_of(%v.map)`, maps.Element()).Location()
}
