package Generate

import "github.com/martinre2/Evolution/Generation"

type GenPoints interface {
	Generate(numer int) Generation.Generation
	Generar() bool
	DoPoints(index int) int
}
