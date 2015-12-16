package Board

import (
	"fmt"
	c "github.com/martinre2/Evolution/Conformation"
)

type BlackBoardParameters struct {
	LatticeType       int
	OptimizationAlg   int
	SelectionOp       int
	CrossOverOp       int
	MutationOp        int
	OtherTechs        []int
	NumConformations  int
	NumGenerations    int
	NumExperiments    int
	Sampling          int
	PercentTournamen  int
	PercentTopPercent int
	CrossOverProb     float32
	MinimumMutation   float32
	MaximumMutation   float32
	ProximityPairing  int
	RandomAssurance   int
	ActualExperiment  int
	BestFitness       int
	KnowBestFitness   bool
}

type BlackBoard struct {
	Params      BlackBoardParameters
	Hplevel     HPLevel
	Generations []Generation
	Bests       []c.Conformation
	BestFitness int
}

func NewBoard(parameters BlackBoardParameters) *BlackBoard {
	fmt.Println(" ")
	board := new(BlackBoard)
	board.Params = parameters
	return board
}

func (b *BlackBoard) IsAvailable(puntosHijo *[]c.Point, puntosHijo_C []c.Point, x float32, y float32, z float32) bool {

	for _, p := range *puntosHijo {
		if p.XValue == x && p.YValue == y && p.ZValue == z {
			return false
		}
	}

	for _, p := range puntosHijo_C {
		if p.XValue == x && p.YValue == y && p.ZValue == z {
			return false
		}
	}
	return true
}

func (b *BlackBoard) GenDmaxP(conformations []c.Conformation) float64 {
	var DmaxP float64 = 0
	for _, c := range conformations {
		DmaxP += c.Dmax
	}
	DmaxP /= float64(len(conformations))
	return DmaxP
}

func (b *BlackBoard) GenRadioGiroP(conformations []c.Conformation) float64 {
	var radGirP float64 = 0
	for _, c := range conformations {
		radGirP += c.RadioGiroRCDM
	}
	radGirP /= float64(len(conformations))
	return radGirP
}

func (b *BlackBoard) AddBest(conformation c.Conformation) {
	b.Bests = append(b.Bests, conformation)
}
