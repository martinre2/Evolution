package Generate

import "github.com/martinre2/Evolution/Board"

type GenPoints interface {
	Generate(numer int) Board.Generation
	Generar() bool
	DoPoints(index int) int
}
