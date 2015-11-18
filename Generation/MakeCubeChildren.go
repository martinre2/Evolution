package Generation

import (
	"fmt"
	b "github.com/martinre2/Evolution/Board"
	c "github.com/martinre2/Evolution/Conformation"
	f "github.com/martinre2/Evolution/Fitness"
	"math/rand"
	_ "sort"
	"strings"
)

type MakeCubeChildren struct {
	board *b.BlackBoard
}

func (mc MakeCubeChildren) MakeChildren(gen b.Generation, board *b.BlackBoard) b.Generation {
	mc.board = board
	var length int = len(board.Hplevel.HpString)
	var numConf int = len(gen.ParentsLst)

	fmt.Println("MC>", length, numConf)

	var conf []c.Conformation
	var corte []int

	cOver := NewCrossOver(board)

	for i := 0; i < numConf; i++ {
		var tempP c.Parents

		if board.Params.CrossOverOp == 0 {
			corte = cOver.OnePointCrossOver(length)
		} else if board.Params.CrossOverOp == 1 {
			corte = cOver.TwoPointCrossOver(length)
		} else if board.Params.CrossOverOp == 2 {
			corte = cOver.UniformPointCrossOver(length)
		}
		var padre1 int = 0
		var padre2 int = 0

		if board.Params.MutationOp == 0 {
			padre1 = gen.ParentsLst[i].Parent1
			padre2 = gen.ParentsLst[i].Parent2
			tempP = gen.ParentsLst[i]
		} else {
			switch rand.Intn(2) {
			case 0:
				padre1 = gen.ParentsLst[i].Parent1
				padre2 = gen.ParentsLst[i].Parent2
				tempP = gen.ParentsLst[i]
				break
			case 1:
				padre1 = gen.ParentsLst[i].Parent2
				padre2 = gen.ParentsLst[i].Parent1
				tempP.Parent1 = padre1
				tempP.Parent2 = padre2
				break
			default:
				fmt.Println("MakeCubeChildrens")
				break
			}
		}

		var puntosPadre1 []c.Point
		puntosPadre1 = make([]c.Point, len(gen.Conformations[padre1].Conformations))
		copy(puntosPadre1, gen.Conformations[padre1].Conformations)
		var puntosPadre2 []c.Point
		puntosPadre2 = make([]c.Point, len(gen.Conformations[padre2].Conformations))
		copy(puntosPadre2, gen.Conformations[padre2].Conformations)

		var nuevoHijo1 []c.Point
		var nuevoHijo2 []c.Point

		for m := 0; m < length; m++ {
			puntosPadre1[m].ResetR()
			puntosPadre2[m].ResetR()
		}

		var puntosHijo_C []c.Point

		fP := c.Point{
			XValue:       0.0,
			YValue:       0.0,
			ZValue:       0.0,
			MovVectValue: 0,
			Letter:       puntosPadre1[0].Letter,
			Way0:         true,
			Way1:         true,
			Way2:         true,
			Way3:         true,
			Way4:         true,
			Way5:         true,
		}

		nuevoHijo1 = append(nuevoHijo1, fP)

		if board.Params.CrossOverOp == 0 {
			for j := 1; j < length; j++ {
				fmt.Println(" It j", j)
				if j < corte[0] {
					j = mc.check(puntosHijo_C, puntosPadre1[j].MovVectValue, &nuevoHijo1, j)
				} else {
					j = mc.check(puntosHijo_C, puntosPadre2[j].MovVectValue, &nuevoHijo1, j)
				}
			}
			var puntosHijo_C []c.Point

			c2 := c.Point{
				XValue:       0.0,
				YValue:       0.0,
				ZValue:       0.0,
				MovVectValue: 0,
				Letter:       puntosPadre2[0].Letter,
				Way0:         true,
				Way1:         true,
				Way2:         true,
				Way3:         true,
				Way4:         true,
				Way5:         true,
			}
			nuevoHijo2 = append(nuevoHijo2, c2)

			for j := 1; j < length; j++ {
				if j < corte[0] {
					j = mc.check(puntosHijo_C, puntosPadre2[j].MovVectValue, &nuevoHijo2, j)
				} else {
					j = mc.check(puntosHijo_C, puntosPadre1[j].MovVectValue, &nuevoHijo2, j)
				}
			}
		} else if board.Params.CrossOverOp == 1 {
			for j := 1; j < length; j++ {
				if j < corte[0] {
					j = mc.check(puntosHijo_C, puntosPadre1[j].MovVectValue, &nuevoHijo1, j)
				} else if j < corte[1] {
					j = mc.check(puntosHijo_C, puntosPadre2[j].MovVectValue, &nuevoHijo1, j)
				} else {
					j = mc.check(puntosHijo_C, puntosPadre1[j].MovVectValue, &nuevoHijo1, j)
				}
			}

			var puntosHijo_C []c.Point

			c2 := c.Point{
				XValue:       0.0,
				YValue:       0.0,
				ZValue:       0.0,
				MovVectValue: 0,
				Letter:       puntosPadre2[0].Letter,
				Way0:         true,
				Way1:         true,
				Way2:         true,
				Way3:         true,
				Way4:         true,
				Way5:         true,
			}
			nuevoHijo2 = append(nuevoHijo2, c2)

			for j := 1; j < length; j++ {
				if j < corte[0] {
					j = mc.check(puntosHijo_C, puntosPadre2[j].MovVectValue, &nuevoHijo2, j)
				} else if j < corte[1] {
					j = mc.check(puntosHijo_C, puntosPadre1[j].MovVectValue, &nuevoHijo2, j)
				} else {
					j = mc.check(puntosHijo_C, puntosPadre2[j].MovVectValue, &nuevoHijo2, j)
				}
			}
		} else if board.Params.CrossOverOp == 2 {
			for j := 1; j < length; j++ {
				if corte[j] == 0 {
					j = mc.check(puntosHijo_C, puntosPadre1[j].MovVectValue, &nuevoHijo1, j)
				} else {
					j = mc.check(puntosHijo_C, puntosPadre2[j].MovVectValue, &nuevoHijo1, j)
				}
			}
			var puntosHijo_C []c.Point

			c2 := c.Point{
				XValue:       0.0,
				YValue:       0.0,
				ZValue:       0.0,
				MovVectValue: 0,
				Letter:       puntosPadre2[0].Letter,
				Way0:         true,
				Way1:         true,
				Way2:         true,
				Way3:         true,
				Way4:         true,
				Way5:         true,
			}
			nuevoHijo2 = append(nuevoHijo2, c2)

			for j := 1; j < length; j++ {
				if corte[j] == 0 {
					j = mc.check(puntosHijo_C, puntosPadre2[j].MovVectValue, &nuevoHijo2, j)
				} else {
					j = mc.check(puntosHijo_C, puntosPadre1[j].MovVectValue, &nuevoHijo2, j)
				}
			}
		}

		hijo1 := c.Conformation{
			Conformations: nuevoHijo1,
			Fitness:       f.NewFitness(nuevoHijo1).GetFitness(),
			Parents:       tempP,
		}
		hijo1.Init()
		conf = append(conf, hijo1)

		hijo2 := c.Conformation{
			Conformations: nuevoHijo2,
			Fitness:       f.NewFitness(nuevoHijo2).GetFitness(),
			Parents:       tempP,
		}
		hijo2.Init()
		conf = append(conf, hijo2)
	}

	for _, b := range board.Bests {
		conf = append(conf, b)
	}
	board.Bests = board.Bests[:0]

	rGn := b.Generation{
		Conformations: conf,
		DmaxP:         board.GenDmaxP(conf),
		RadioGiroP:    board.GenRadioGiroP(conf),
	}
	return rGn
}

func (mc MakeCubeChildren) check(pointsHijo_C []c.Point, v int, puntosHijo *[]c.Point, j int) int {

	var isOk bool = true
	var sb string = ""
	for {
		pto := mc.genPtoCubico(v, mc.board.Hplevel.HpString[j], (*puntosHijo)[j-1])

		switch rand.Intn(6) {
		case 0:
			if (*puntosHijo)[j-1].Way0 &&
				mc.board.IsAvailable(puntosHijo, pointsHijo_C, pto.XValue, pto.YValue, pto.ZValue) {
				*puntosHijo = append(*puntosHijo, pto)
				isOk = true
			} else {
				isOk = false
				sb += "a"
			}
			break

		case 1:
			if (*puntosHijo)[j-1].Way1 &&
				mc.board.IsAvailable(puntosHijo, pointsHijo_C, pto.XValue, pto.YValue, pto.ZValue) {
				*puntosHijo = append(*puntosHijo, pto)
				isOk = true
			} else {
				isOk = false
				sb += "b"
			}
			break
		case 2:

			if (*puntosHijo)[j-1].Way2 &&
				mc.board.IsAvailable(puntosHijo, pointsHijo_C, pto.XValue, pto.YValue, pto.ZValue) {
				*puntosHijo = append(*puntosHijo, pto)
				isOk = true
			} else {
				isOk = false
				sb += "c"
			}
			break
		case 3:

			if (*puntosHijo)[j-1].Way3 &&
				mc.board.IsAvailable(puntosHijo, pointsHijo_C, pto.XValue, pto.YValue, pto.ZValue) {
				*puntosHijo = append(*puntosHijo, pto)
				isOk = true
			} else {
				isOk = false
				sb += "d"
			}
			break

		case 4:

			if (*puntosHijo)[j-1].Way4 &&
				mc.board.IsAvailable(puntosHijo, pointsHijo_C, pto.XValue, pto.YValue, pto.ZValue) {
				*puntosHijo = append(*puntosHijo, pto)
				isOk = true
			} else {
				isOk = false
				sb += "e"
			}
			break

		case 5:
			if (*puntosHijo)[j-1].Way5 &&
				mc.board.IsAvailable(puntosHijo, pointsHijo_C, pto.XValue, pto.YValue, pto.ZValue) {
				*puntosHijo = append(*puntosHijo, pto)
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

		if !isOk {
			v = rand.Intn(6)
		}

		if strings.Index(sb, "a") != -1 &&
			strings.Index(sb, "b") != -1 &&
			strings.Index(sb, "c") != -1 &&
			strings.Index(sb, "d") != -1 &&
			strings.Index(sb, "e") != -1 &&
			strings.Index(sb, "f") != -1 {

			switch pto.MovVectValue {
			case 0:
				(*puntosHijo)[j-1].Way0 = false
				break
			case 1:
				(*puntosHijo)[j-1].Way1 = false
				break
			case 2:
				(*puntosHijo)[j-1].Way2 = false
				break
			case 3:
				(*puntosHijo)[j-1].Way3 = false
				break
			case 4:
				(*puntosHijo)[j-1].Way4 = false
				break
			case 5:
				(*puntosHijo)[j-1].Way5 = false
				break
			default:
				fmt.Println("GenCubePoints")
				break
			}

			pointsHijo_C = append(pointsHijo_C, (*puntosHijo)[j-1])
			*puntosHijo = append((*puntosHijo)[:j-2], (*puntosHijo)[j-1:]...)
			sb = ""
			j--
		}

		if isOk {
			break
		}

	}

	return j
}

func (mc MakeCubeChildren) genPtoCubico(vectMov int, l byte, ptoAnterior c.Point) c.Point {
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
