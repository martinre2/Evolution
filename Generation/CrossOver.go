package Generation

import (
	_ "fmt"
	b "github.com/martinre2/Evolution/Board"
	"math/rand"
)

type CrossOver struct {
	board *b.BlackBoard
}

func NewCrossOver(b *b.BlackBoard) *CrossOver {
	c := new(CrossOver)
	c.board = b
	return c
}

func (c *CrossOver) OnePointCrossOver(longi int) []int {
	var corte []int
	if c.board.Params.CrossOverProb > rand.Float32() {
		corte = append(corte, rand.Intn(longi-1)+1)
		return corte
	} else {
		corte = append(corte, longi)
		return corte
	}
}

func (c *CrossOver) TwoPointCrossOver(longi int) []int {
	var corte []int
	if c.board.Params.CrossOverProb > rand.Float32() {
		corte[0] = rand.Intn(longi / 2)
		corte[1] = rand.Intn(longi/2) + (longi / 2)
	} else {
		corte[1], corte[0] = longi, longi
	}
	return corte
}

func (c *CrossOver) UniformPointCrossOver(longi int) []int {
	var corte []int
	corte = make([]int, longi, longi)
	if c.board.Params.CrossOverProb > rand.Float32() {
		for i := 0; i < longi; i++ {
			corte[i] = rand.Intn(2)
		}
	}
	return corte
}
