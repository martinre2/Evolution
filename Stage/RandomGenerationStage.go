package Stage

import (
	b "github.com/martinre2/Evolution/Board"
	gn "github.com/martinre2/Evolution/Generate"
)

type RandomGenerationStage struct{}

func (r RandomGenerationStage) Execute(board *b.BlackBoard) bool {
	numConf := board.Params.NumConformations
	//lattType := board.Params.LatticeType
	if numConf%2 != 0 {
		return false
	}

	grdm := gn.NewGenCubePoints(board)

	board.Generations = append(board.Generations, grdm.Generate(numConf))
	return false
}

func (r RandomGenerationStage) String() string { return "RandomGenerationStage" }
