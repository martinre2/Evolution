package Conformation

import (
	"math"
)

type Conformation struct {
	Fitness       int
	Conformations []Point
	Parents       Parents
	RadioGiroCM   [3]float64
	RadioGiroRCDM float64
	Dmax          float64
}

func NewConformation(conformation []Point) *Conformation {
	conf := new(Conformation)
	conf.Conformations = conformation
	conf.genDmax()
	conf.genRadGirCM()
	conf.genRadGirRCDM()
	return conf
}

func (c *Conformation) Init() {
	c.genDmax()
	c.genRadGirCM()
	c.genRadGirRCDM()
}

func (c *Conformation) genRadGirCM() {
	var rmeanX, rmeanY, rmeanZ float32 = 0, 0, 0
	for _, p := range c.Conformations {
		rmeanX += p.XValue
		rmeanY += p.YValue
		rmeanZ += p.ZValue
	}

	rmeanX = rmeanX / float32(len(c.Conformations))
	rmeanY = rmeanY / float32(len(c.Conformations))
	rmeanZ = rmeanZ / float32(len(c.Conformations))

	var tempX, tempY, tempZ float64 = 0, 0, 0
	for _, p := range c.Conformations {
		tempX += math.Pow(float64(p.XValue-rmeanX), 2)
		tempY += math.Pow(float64(p.YValue-rmeanY), 2)
		tempZ += math.Pow(float64(p.ZValue-rmeanZ), 2)
	}
	c.RadioGiroCM[0] = math.Sqrt(tempX / float64(len(c.Conformations)))
	c.RadioGiroCM[1] = math.Sqrt(tempY / float64(len(c.Conformations)))
	c.RadioGiroCM[2] = math.Sqrt(tempZ / float64(len(c.Conformations)))
}

func (c *Conformation) genRadGirRCDM() {
	var N int = len(c.Conformations)
	var temp float64 = 0

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			temp += math.Pow(float64(c.Conformations[i].XValue-c.Conformations[j].XValue), 2) +
				math.Pow(float64(c.Conformations[i].YValue-c.Conformations[j].YValue), 2) +
				math.Pow(float64(c.Conformations[i].ZValue-c.Conformations[j].ZValue), 2)
		}
	}
	temp /= math.Pow(float64(N), 2)
	c.RadioGiroRCDM = math.Sqrt(temp)
}

func (c *Conformation) genDmax() {

	var N int = len(c.Conformations)
	var Dmax, temp float64 = 0, 0
	for i := 0; i < N; i++ {
		temp = 0
		for j := i + 1; j < N; j++ {
			temp = math.Sqrt(math.Pow(float64(c.Conformations[i].XValue-c.Conformations[j].XValue), 2) + math.Pow(float64(c.Conformations[i].YValue-c.Conformations[j].YValue), 2) + math.Pow(float64(c.Conformations[i].ZValue-c.Conformations[j].ZValue), 2))
			if Dmax < temp {
				Dmax = temp
			}
		}
	}
	c.Dmax = Dmax
}

type Conformations []Conformation

func (slice Conformations) Len() int {
	return len(slice)
}

func (slice Conformations) Less(i, j int) bool {
	return slice[i].Fitness < slice[j].Fitness
}

func (slice Conformations) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
