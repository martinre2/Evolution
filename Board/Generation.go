package Board

import (
	c "github.com/martinre2/Evolution/Conformation"
)

type Generation struct {
	Conformations []c.Conformation
	Selection     []int
	ParentsLst    []c.Parents
	RadioGiroP    float64
	DmaxP         float64
}
