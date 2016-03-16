package Selection

import (
	"fmt"
	b "github.com/martinre2/Evolution/Board"
	c "github.com/martinre2/Evolution/Conformation"
	"math/rand"
	"sort"
)

type Roulette struct{}

func (r Roulette) Roulette(generation *b.Generation, board *b.BlackBoard) bool {

	//fmt.Println("Gen", generation)
	var ln int = len(generation.Conformations)
	var totalFitness int = 0
	for i := 0; i < ln; i++ {
		totalFitness += generation.Conformations[i].Fitness
	}

	fmt.Println("TotFit", totalFitness)

	totalFitness *= -1

	if totalFitness == 0 {
		return true
	}

	var seleccion []int

	if board.Params.OtherTechs[0] == 1 {
		var dest c.Conformations
		dest = make([]c.Conformation, len(generation.Conformations))
		copy(dest, generation.Conformations)

		sort.Sort(dest)
		//fmt.Println("____>>dest", dest[0])
		fmt.Println("____>>dest0", dest[0])
		fmt.Println("____>>dest1", dest[1])
		board.AddBest(dest[0])
		board.AddBest(dest[1])

		//fmt.Println("BOARD", board.Bests)

		dest = dest[:0]

		for len(seleccion) < ln-2 {
			valor := rand.Float32()
			valor *= float32(totalFitness)
			var partialSum float32 = 0

			for i := 0; i < ln; i++ {
				if partialSum <= valor {
					partialSum += float32(generation.Conformations[i].Fitness * -1)
				} else {
					if r.indexOf(seleccion, i-1) == r.lastIndexOf(seleccion, i-1) {
						seleccion = append(seleccion, i-1)
						break
					}
				}
			}
		}

		fmt.Println("____>>SElECT", seleccion)
	} else if board.Params.OtherTechs[0] == 2 {
		/*ArrayList<Conformation> dest = (ArrayList<Conformation>) generation.getConformations().clone();
		board.Params.ordenar(dest);
		board.Params.addBest(dest.get(0));
		dest.clear();
		dest = null;
		while (seleccion.size() < ln - 1) {
			double valor = rand.nextDouble();
			valor *= totalFitness;
			double partialSum = 0;
			for (int i = 0; i < ln; i++) {
				if (partialSum <= valor) {
					partialSum += generation.getConformations().get(i).getFitness() * -1;
				} else {
					if (seleccion.indexOf(i - 1) == seleccion.lastIndexOf(i - 1)) {
						seleccion.add(i - 1);
						break;
					}
				}
			}
		}*/
	} else {
		/*while (seleccion.size() < ln) {
			double valor = rand.nextDouble();
			valor *= totalFitness;
			double partialSum = 0;
			for (int i = 0; i < ln; i++) {
				if (partialSum <= valor) {
					partialSum += generation.getConformations().get(i).getFitness() * -1;
				} else {
					if (seleccion.indexOf(i - 1) == seleccion.lastIndexOf(i - 1)) {
						seleccion.add(i - 1);
						break;
					}
				}
			}
		}*/
	}
	sort.Ints(seleccion)
	generation.Selection = make([]int, len(seleccion))
	copy(generation.Selection, seleccion)
	fmt.Println("____>>SELECTION", generation.Selection)
	return false
}

func (r Roulette) indexOf(arr []int, s int) int {
	for i, e := range arr {
		if e == s {
			return i
		}
	}
	return -1
}

func (r Roulette) lastIndexOf(arr []int, s int) int {
	ret := -1
	for i, e := range arr {
		if e == s {
			ret = i
		}
	}
	return ret
}
