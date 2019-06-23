package maps

type Location [2]float64

type Options struct {
	Center Location `json:"center,omitempty"`
	Zoom float64 `json:"zoom,omitempty"`
}
