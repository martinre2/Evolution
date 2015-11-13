package Generate

import (
	"fmt"
	b "github.com/martinre2/Evolution/Board"
	c "github.com/martinre2/Evolution/Conformation"
	f "github.com/martinre2/Evolution/Fitness"
	gn "github.com/martinre2/Evolution/Generation"
	"math/rand"
	"strings"
)

type GenCubePoints struct {
	HPsec  string
	HpLen  int
	Points []c.Point
	board  b.BlackBoard
}

func NewGenCubePoints(board *b.BlackBoard) *GenCubePoints {
	genRdm := new(GenCubePoints)
	genRdm.HPsec = board.Hplevel.HpString
	genRdm.HpLen = len(board.Hplevel.HpString)
	board = board

	return genRdm
}

func (g *GenCubePoints) Generate(number int) gn.Generation {
	var conformaciones []c.Conformation
	fmt.Println("Conformations Generator", number)
	for i := 0; i < number; i++ {

		fmt.Println("--> conf:", i)
		var puntosHijo_C []c.Point
		g.Points = g.Points[:0]

		g.generar(puntosHijo_C)
		fitness := f.NewFitness(g.Points).GetFitness()

		if fitness == 0 {
			i = i - 1
			continue
		}

		conf := c.Conformation{
			Conformations: g.Points,
			Fitness:       f.NewFitness(g.Points).GetFitness(),
		}
		conf.Init()

		conformaciones = append(conformaciones, conf)
	}

	rGn := gn.Generation{
		Conformations: conformaciones,
		DmaxP:         g.board.GenDmaxP(conformaciones),
		RadioGiroP:    g.board.GenRadioGiroP(conformaciones),
	}
	return rGn
}

func (g *GenCubePoints) generar(puntosHijo_C []c.Point) {

	fP := c.Point{
		XValue:       0.0,
		YValue:       0.0,
		ZValue:       0.0,
		MovVectValue: 0,
		Letter:       g.HPsec[0],
		Way0:         true,
		Way1:         true,
		Way2:         true,
		Way3:         true,
		Way4:         true,
		Way5:         true,
	}

	g.Points = append(g.Points, fP)
	for i := 1; i < g.HpLen; i++ {
		i = g.doPoints(puntosHijo_C, i)
	}
}

func (g *GenCubePoints) doPoints(pointsHijo_C []c.Point, i int) int {

	var isOk bool = true
	var sb string = ""
	for {

		switch rand.Intn(6) {
		case 0:

			if g.Points[i-1].Way0 &&
				g.board.IsAvailable(g.Points, pointsHijo_C, g.Points[i-1].XValue+1, g.Points[i-1].YValue, g.Points[i-1].ZValue) {
				g.Points = append(g.Points, g.genPtoCubico(0, g.HPsec[i], g.Points[i-1]))
				isOk = true
			} else {
				isOk = false
				sb += "a"
			}
			break

		case 1:
			if g.Points[i-1].Way0 &&
				g.board.IsAvailable(g.Points, pointsHijo_C, g.Points[i-1].XValue+1, g.Points[i-1].YValue, g.Points[i-1].ZValue) {
				g.Points = append(g.Points, g.genPtoCubico(1, g.HPsec[i], g.Points[i-1]))
				isOk = true
			} else {
				isOk = false
				sb += "b"
			}
			break
		case 2:

			if g.Points[i-1].Way0 &&
				g.board.IsAvailable(g.Points, pointsHijo_C, g.Points[i-1].XValue+1, g.Points[i-1].YValue, g.Points[i-1].ZValue) {
				g.Points = append(g.Points, g.genPtoCubico(2, g.HPsec[i], g.Points[i-1]))
				isOk = true
			} else {
				isOk = false
				sb += "c"
			}
			break
		case 3:

			if g.Points[i-1].Way0 &&
				g.board.IsAvailable(g.Points, pointsHijo_C, g.Points[i-1].XValue+1, g.Points[i-1].YValue, g.Points[i-1].ZValue) {
				g.Points = append(g.Points, g.genPtoCubico(3, g.HPsec[i], g.Points[i-1]))
				isOk = true
			} else {
				isOk = false
				sb += "d"
			}
			break

		case 4:

			if g.Points[i-1].Way0 &&
				g.board.IsAvailable(g.Points, pointsHijo_C, g.Points[i-1].XValue+1, g.Points[i-1].YValue, g.Points[i-1].ZValue) {
				g.Points = append(g.Points, g.genPtoCubico(4, g.HPsec[i], g.Points[i-1]))
				isOk = true
			} else {
				isOk = false
				sb += "e"
			}
			break

		case 5:
			if g.Points[i-1].Way0 &&
				g.board.IsAvailable(g.Points, pointsHijo_C, g.Points[i-1].XValue+1, g.Points[i-1].YValue, g.Points[i-1].ZValue) {
				g.Points = append(g.Points, g.genPtoCubico(5, g.HPsec[i], g.Points[i-1]))
				isOk = true
			} else {
				isOk = false
				sb += "f"
			}
			break
		default:
			fmt.Println("GenCubePoints")
			break
		}

		if strings.Index(sb, "a") != -1 &&
			strings.Index(sb, "b") != -1 &&
			strings.Index(sb, "c") != -1 &&
			strings.Index(sb, "d") != -1 &&
			strings.Index(sb, "e") != -1 &&
			strings.Index(sb, "f") != -1 {

			switch g.Points[i-1].MovVectValue {
			case 0:
				g.Points[i-1].Way0 = false
				break
			case 1:
				g.Points[i-1].Way1 = false
				break
			case 2:
				g.Points[i-1].Way2 = false
				break
			case 3:
				g.Points[i-1].Way3 = false
				break
			case 4:
				g.Points[i-1].Way4 = false
				break
			case 5:
				g.Points[i-1].Way5 = false
				break
			default:
				fmt.Println("GenCubePoints")
				break
			}

			pointsHijo_C = append(pointsHijo_C, g.Points[i-1])
			g.Points = append(g.Points[:i-1], g.Points[i:]...)
			sb = ""
			i--
		}

		if isOk {
			break
		}

	}
	return i
}

func (g *GenCubePoints) genPtoCubico(vectMov int, l byte, ptoAnterior c.Point) c.Point {
	p := c.Point{
		XValue:       0.0,
		YValue:       0.0,
		ZValue:       0.0,
		MovVectValue: 0,
		Letter:       l,
		Way0:         true,
		Way1:         true,
		Way2:         true,
		Way3:         true,
		Way4:         true,
		Way5:         true,
	}
	switch vectMov {
	case 0:
		p.XValue = ptoAnterior.XValue + 1
		p.YValue = ptoAnterior.YValue
		p.ZValue = ptoAnterior.ZValue
		p.MovVectValue = 0
		p.Letter = l
		p.Way1 = false
		break
	case 1:
		p.XValue = ptoAnterior.XValue - 1
		p.YValue = ptoAnterior.YValue
		p.ZValue = ptoAnterior.ZValue
		p.MovVectValue = 1
		p.Letter = l
		p.Way0 = false
		break
	case 2:
		p.XValue = ptoAnterior.XValue
		p.YValue = ptoAnterior.YValue + 1
		p.ZValue = ptoAnterior.ZValue
		p.MovVectValue = 2
		p.Letter = l
		p.Way3 = false
		break
	case 3:
		p.XValue = ptoAnterior.XValue
		p.YValue = ptoAnterior.YValue - 1
		p.ZValue = ptoAnterior.ZValue
		p.MovVectValue = 3
		p.Letter = l
		p.Way2 = false
		break
	case 4:
		p.XValue = ptoAnterior.XValue
		p.YValue = ptoAnterior.YValue
		p.ZValue = ptoAnterior.ZValue + 1
		p.MovVectValue = 4
		p.Letter = l
		p.Way5 = false
		break
	case 5:
		p.XValue = ptoAnterior.XValue
		p.YValue = ptoAnterior.YValue
		p.ZValue = ptoAnterior.ZValue - 1
		p.MovVectValue = 5
		p.Letter = l
		p.Way4 = false
		break
	default:
		fmt.Println("GenCubePoints")
		break
	}
	return p
}
