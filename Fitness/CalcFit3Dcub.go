package Fitness

import (
	c "github.com/martinre2/Evolution/Conformation"
)

type CalcFit3Dcub struct {
	Points   []c.Point
	fitCount int
	posAct   int
	posSig   int
	isHidro  bool
}

func NewCalcFit3Dcub(points []c.Point) *CalcFit3Dcub {
	cF := new(CalcFit3Dcub)
	cF.Points = points
	return cF
}

func (f *CalcFit3Dcub) get3dFitness() int {
	f.fitCount = 0
	var pSize int = len(f.Points)
	// Check FIRST point
	var punto c.Point = f.Points[0]

	f.posSig = f.Points[1].MovVectValue

	if string(punto.Letter) == "H" {
		if f.posSig == 0 {
			// check Back
			if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Up
			if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
				f.fitCount--
			}
			// check Down
			if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
				f.fitCount--
			}
			// check Right
			if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
				f.fitCount--
			}
			// check Left
			if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
				f.fitCount--
			}
		} else if f.posSig == 1 {
			// check Forward
			if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Up
			if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
				f.fitCount--
			}
			// check Down
			if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
				f.fitCount--
			}
			// check Right
			if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
				f.fitCount--
			}
			// check Left
			if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
				f.fitCount--
			}
		} else if f.posSig == 2 {
			// check Forward
			if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Back
			if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Down
			if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
				f.fitCount--
			}
			// check Right
			if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
				f.fitCount--
			}
			// check Left
			if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
				f.fitCount--
			}
		} else if f.posSig == 3 {
			// check Forward
			if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Back
			if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Up
			if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
				f.fitCount--
			}
			// check Right
			if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
				f.fitCount--
			}
			// check Left
			if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
				f.fitCount--
			}
		} else if f.posSig == 4 {
			// check Forward
			if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Back
			if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Up
			if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
				f.fitCount--
			}
			// check Down
			if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
				f.fitCount--
			}
			// check Left
			if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
				f.fitCount--
			}
		} else if f.posSig == 5 {
			// check Forward
			if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Back
			if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Up
			if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
				f.fitCount--
			}
			// check Down
			if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
				f.fitCount--
			}
			// check Right
			if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
				f.fitCount--
			}
		}
	}

	// Check LAST point
	punto = f.Points[pSize-1]
	f.posAct = punto.MovVectValue
	if punto.Letter == 'H' {
		if f.posAct == 0 {
			// check Forward
			if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Up
			if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
				f.fitCount--
			}
			// check Down
			if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
				f.fitCount--
			}
			// check Right
			if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
				f.fitCount--
			}
			// check Left
			if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
				f.fitCount--
			}
		} else if f.posAct == 1 {
			// check Back
			if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Up
			if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
				f.fitCount--
			}
			// check Down
			if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
				f.fitCount--
			}
			// check Right
			if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
				f.fitCount--
			}
			// check Left
			if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
				f.fitCount--
			}
		} else if f.posAct == 2 {
			// check Forward
			if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Back
			if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Up
			if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
				f.fitCount--
			}
			// check Right
			if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
				f.fitCount--
			}
			// check Left
			if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
				f.fitCount--
			}
		} else if f.posAct == 3 {
			// check Forward
			if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Back
			if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Down
			if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
				f.fitCount--
			}
			// check Right
			if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
				f.fitCount--
			}
			// check Left
			if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
				f.fitCount--
			}
		} else if f.posAct == 4 {
			// check Forward
			if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Back
			if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Up
			if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
				f.fitCount--
			}
			// check Down
			if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
				f.fitCount--
			}
			// check Right
			if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
				f.fitCount--
			}
		} else if f.posAct == 5 {
			// check Forward
			if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Back
			if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
				f.fitCount--
			}
			// check Up
			if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
				f.fitCount--
			}
			// check Down
			if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
				f.fitCount--
			}
			// check Left
			if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
				f.fitCount--
			}
		}
	}

	// Check Middle points
	for i := 1; i < pSize-1; i++ {
		punto = f.Points[i]
		if punto.Letter == 'H' {
			f.posAct = punto.MovVectValue
			f.posSig = f.Points[i+1].MovVectValue
			if f.posAct == 0 {
				if f.posSig == 0 {
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 2 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 3 {
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 4 {
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 5 {
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
				}
			} else if f.posAct == 1 {
				if f.posSig == 1 {
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 2 {
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 3 {
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 4 {
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 5 {
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
				}
			} else if f.posAct == 2 {
				if f.posSig == 0 {
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 1 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 2 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 4 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 5 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
				}
			} else if f.posAct == 3 {
				if f.posSig == 0 {
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 1 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 3 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 4 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 5 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
				}
			} else if f.posAct == 4 {
				if f.posSig == 0 {
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
				} else if f.posSig == 1 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
				} else if f.posSig == 2 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
				} else if f.posSig == 3 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Right
					if f.isH(punto.XValue, punto.YValue, punto.ZValue+1) {
						f.fitCount--
					}
				} else if f.posSig == 4 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
				}
			} else if f.posAct == 5 {
				if f.posSig == 0 {
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 1 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 2 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 3 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Left
					if f.isH(punto.XValue, punto.YValue, punto.ZValue-1) {
						f.fitCount--
					}
				} else if f.posSig == 5 {
					// check Forward
					if f.isH(punto.XValue+1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Back
					if f.isH(punto.XValue-1, punto.YValue, punto.ZValue) {
						f.fitCount--
					}
					// check Up
					if f.isH(punto.XValue, punto.YValue+1, punto.ZValue) {
						f.fitCount--
					}
					// check Down
					if f.isH(punto.XValue, punto.YValue-1, punto.ZValue) {
						f.fitCount--
					}
				}
			}
		}
	}

	return (f.fitCount / 2)
}

func (f *CalcFit3Dcub) isH(xValue float32, yValue float32, zValue float32) bool {
	var isHidro bool = false

	for _, p := range f.Points {
		if p.XValue == xValue &&
			p.YValue == yValue &&
			p.ZValue == zValue &&
			string(p.Letter) == "H" {
			isHidro = true
			break
		}
	}
	return isHidro
}
