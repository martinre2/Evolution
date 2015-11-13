package Stage

import (
	b "github.com/martinre2/Evolution/Board"
	cf "github.com/martinre2/Evolution/CoupleFormation"
	s "github.com/martinre2/Evolution/Selection"
	//gn "github.com/martinre2/Evolution/Generate"
	"fmt"
)

type SimpleGenerationStage struct {
	board *b.BlackBoard
}

func (r SimpleGenerationStage) Execute(board *b.BlackBoard) bool {
	r.board = board
	var ex bool = false
	if board.Params.KnowBestFitness {
		ex = r.runV2()
	} else {
		ex = r.runV1()
	}
	return ex
}

func (r SimpleGenerationStage) runV1() bool {
	var numGene int = r.board.Params.NumGenerations
	var sampling int = r.board.Params.Sampling
	_ = sampling
	var operSel int = r.board.Params.SelectionOp
	var actExp int = r.board.Params.ActualExperiment
	_ = actExp
	var percent int = 0
	_ = percent
	var j int = 1
	_ = j

	fmt.Println("SimpleGen>", numGene)
	for i := 0; i < numGene; i++ {
		fmt.Println("SimpleGen>", i)
		if operSel == 0 {
			if new(s.Roulette).Roulette(r.board.Generations[i], r.board) {
				return false
			}
		} else if operSel == 1 {
			//percent = r.board.getPercentTournamen();
			//if (new Tournament().tournament(percent, r.board.getProjLvl().getExperiment(actExp).get(i))) {
			//	return false;
			//}
		} else if operSel == 2 {
			//percent = r.board.getPercentTopPercent();
			//if (new TopPercent().percent(percent, r.board.getProjLvl().getExperiment(actExp).get(i))) {
			//	return false;
			//}
		} else if operSel == 3 {
			//if (new PopulDecim().population(r.board.getProjLvl().getExperiment(actExp).get(i))) {
			//return false;
			//}
		}

		new(cf.SimplexCoupleFormation).CoupleFormation(r.board.Generations[i], r.board)
		fmt.Println("PARENTS GEN", r.board.Generations[i].ParentsLst)

		if r.board.getLatticeType() == 0 {
			//r.board.Generations = append(r.board.Generations, null)
		} else if r.board.getLatticeType() == 1 {
			//r.board.Generations = append(r.board.Generations, null)
		} else if r.board.getLatticeType() == 2 {
			r.board.Generations = append(r.board.Generations, new(MakeCubeChildren).MakeChildren(r.board.Generations[i], r.board))
			//r.board.getProjLvl().getExperiment(actExp).add(new MakeCubeChildren().makeChildren(r.board.getProjLvl().getExperiment(actExp).get(i)));
		}
		/*
			if ((i % sampling) == 0) {
				for (int k = j; k < i; k++) {
					r.board.getProjLvl().getExperiment(actExp).set(k, null);
				}
				j = i + 1;
			}
		*/
	}
	/*
		int b = r.board.getProjLvl().getExperiment(actExp).size();
		for (int k = j; k < b - 1; k++) {
			r.board.getProjLvl().getExperiment(actExp).set(k, null);
		}
		for (int i = 0; i < b; i++) {
			r.board.getProjLvl().getExperiment(actExp).remove(null);
		}
	*/
	return true
}

func (r SimpleGenerationStage) runV2() bool {
	return true
}
