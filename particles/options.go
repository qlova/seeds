package particles

type Options struct {
	Particles     Particles     `json:"particles"`
	Interactivity Interactivity `json:"interactivity"`
	RetinaDetect  bool          `json:"retina_detect"`
}
type Density struct {
	Enable    bool `json:"enable"`
	ValueArea int  `json:"value_area"`
}
type Number struct {
	Value   int     `json:"value"`
	Density Density `json:"density"`
}
type Color struct {
	Value string `json:"value"`
}
type Stroke struct {
	Width int    `json:"width"`
	Color string `json:"color"`
}
type Polygon struct {
	NbSides int `json:"nb_sides"`
}
type Image struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
type Shape struct {
	Type    string  `json:"type"`
	Stroke  Stroke  `json:"stroke"`
	Polygon Polygon `json:"polygon"`
	Image   Image   `json:"image"`
}
type Anim struct {
	Enable     bool    `json:"enable"`
	Speed      int     `json:"speed"`
	OpacityMin int     `json:"opacity_min"`
	SizeMin    float64 `json:"size_min"`
	Sync       bool    `json:"sync"`
}
type Opacity struct {
	Value  int  `json:"value"`
	Random bool `json:"random"`
	Anim   Anim `json:"anim"`
}
type Size struct {
	Value  int  `json:"value"`
	Random bool `json:"random"`
	Anim   Anim `json:"anim"`
}
type LineLinked struct {
	Enable   bool    `json:"enable"`
	Distance int     `json:"distance"`
	Color    string  `json:"color"`
	Opacity  float64 `json:"opacity"`
	Width    int     `json:"width"`
}
type Attract struct {
	Enable  bool `json:"enable"`
	RotateX int  `json:"rotateX"`
	RotateY int  `json:"rotateY"`
}
type Move struct {
	Enable    bool    `json:"enable"`
	Speed     int     `json:"speed"`
	Direction string  `json:"direction"`
	Random    bool    `json:"random"`
	Straight  bool    `json:"straight"`
	OutMode   string  `json:"out_mode"`
	Bounce    bool    `json:"bounce"`
	Attract   Attract `json:"attract"`
}
type Particles struct {
	Number     Number     `json:"number"`
	Color      Color      `json:"color"`
	Shape      Shape      `json:"shape"`
	Opacity    Opacity    `json:"opacity"`
	Size       Size       `json:"size"`
	LineLinked LineLinked `json:"line_linked"`
	Move       Move       `json:"move"`
}
type Onhover struct {
	Enable bool   `json:"enable"`
	Mode   string `json:"mode"`
}
type Onclick struct {
	Enable bool   `json:"enable"`
	Mode   string `json:"mode"`
}
type Events struct {
	Onhover Onhover `json:"onhover"`
	Onclick Onclick `json:"onclick"`
	Resize  bool    `json:"resize"`
}
type Grab struct {
	Distance   int        `json:"distance"`
	LineLinked LineLinked `json:"line_linked"`
}
type Bubble struct {
	Distance int `json:"distance"`
	Size     int `json:"size"`
	Duration int `json:"duration"`
	Opacity  int `json:"opacity"`
	Speed    int `json:"speed"`
}
type Repulse struct {
	Distance int     `json:"distance"`
	Duration float64 `json:"duration"`
}
type Push struct {
	ParticlesNb int `json:"particles_nb"`
}
type Remove struct {
	ParticlesNb int `json:"particles_nb"`
}
type Modes struct {
	Grab    Grab    `json:"grab"`
	Bubble  Bubble  `json:"bubble"`
	Repulse Repulse `json:"repulse"`
	Push    Push    `json:"push"`
	Remove  Remove  `json:"remove"`
}
type Interactivity struct {
	DetectOn string `json:"detect_on"`
	Events   Events `json:"events"`
	Modes    Modes  `json:"modes"`
}
