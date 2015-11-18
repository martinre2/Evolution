package main

import (
	"fmt"
	b "github.com/martinre2/Evolution/Board"
	r "github.com/martinre2/Evolution/Stage"
)

func main() {
	params := b.BlackBoardParameters{
		LatticeType:      2,
		NumConformations: 18,
		NumGenerations:   100,
		NumExperiments:   1,
		OtherTechs:       []int{1, 0, 0},
		CrossOverProb:    1.0,
	}

	board := b.NewBoard(params)
	board.Hplevel.HpString = "HHHHHHHHPPPPHHHHH"

	ex_randomStage := new(r.RandomGenerationStage).Execute(board)
	ex_simpleStage := new(r.SimpleGenerationStage).Execute(board)

	_ = board
	_ = ex_randomStage
	_ = ex_simpleStage
	fmt.Println(len(board.Generations))
	fmt.Println(ex_randomStage)
	fmt.Println(ex_simpleStage)
}
