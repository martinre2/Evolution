package Fitness

import (
	c "github.com/martinre2/Evolution/Conformation"
)

type Fitness struct {
	Points []c.Point
}

func NewFitness(points []c.Point) *Fitness {
	f := new(Fitness)
	f.Points = points
	return f
}

func (f *Fitness) GetFitness() int {
	calc := NewCalcFit3Dcub(f.Points)
	return calc.get3dFitness()
}
