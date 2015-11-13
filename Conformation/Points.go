package Conformation

type Point struct {
	XValue       float32
	YValue       float32
	ZValue       float32
	MovVectValue int
	Letter       byte
	Way0         bool
	Way1         bool
	Way2         bool
	Way3         bool
	Way4         bool
	Way5         bool
}

func NewPoint() *Point {
	p := new(Point)
	p.XValue = 0.0
	p.YValue = 0.0
	p.ZValue = 0.0
	p.MovVectValue = 0
	p.Letter = ' '
	p.Way0 = true
	p.Way1 = true
	p.Way2 = true
	p.Way3 = true
	p.Way4 = true
	p.Way5 = true

	return p
}

func (p *Point) ResetR() {
	p.Way0 = true
	p.Way1 = true
	p.Way2 = true
	p.Way3 = true
	p.Way4 = true
	p.Way5 = true
}
