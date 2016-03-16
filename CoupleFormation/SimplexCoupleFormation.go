package CoupleFormation

import (
	"fmt"
	b "github.com/martinre2/Evolution/Board"
	c "github.com/martinre2/Evolution/Conformation"
	"math/rand"
)

type SimplexCoupleFormation struct{}

func (r SimplexCoupleFormation) CoupleFormation(generation *b.Generation, board *b.BlackBoard) {
	selRuleta := generation.Selection
	var padres []c.Parents
	var part2 []int
	var tam int = len(selRuleta) / 2

	for i := 0; i < tam; i++ {
		padres = append(padres, c.Parents{Parent1: selRuleta[i]})
		part2 = append(part2, selRuleta[i+tam])
	}

	fmt.Println("Selector", selRuleta)
	fmt.Println("padres", padres)
	fmt.Println("part2", part2)

	for i := 0; i < tam; i++ {
		var rdm int = rand.Intn(len(part2))

		p := c.Parents{
			Parent1: padres[i].Parent1,
			Parent2: part2[rdm],
		}
		var bol bool = true

		for _, pa := range padres {
			if pa.Parent1 == p.Parent1 &&
				pa.Parent2 == p.Parent2 {
				if len(part2) == 1 {
					bol = true
					break
				} else {
					i--
					bol = false
					break
				}
			}
		}
		if bol {
			padres[i].Parent2 = part2[rdm]
			fmt.Println("part2", part2)
			fmt.Println("rdm", rdm)
			part2 = append(part2[:rdm], part2[rdm+1:]...)
			fmt.Println("part2", part2)
		}
	}
	generation.ParentsLst = make([]c.Parents, len(padres))
	fmt.Println("part2", padres)
	copy(generation.ParentsLst, padres)

}
