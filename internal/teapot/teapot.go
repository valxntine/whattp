package teapot

type Teapot struct {
	LineOne   string
	LineTwo   string
	LineThree string
	LineFour  string
	LineFive  string
	LineSix   string
}

func NewTeapot() *Teapot {
	return &Teapot{
		LineOne:   "                     ;,'",
		LineTwo:   "           _o_      ;:;'",
		LineThree: "   ,-.'---'.__ ;",
		LineFour:  " ((j;=====',-'",
		LineFive:  "'-\\     /",
		LineSix:   "      '-=-'     ",
	}
}
