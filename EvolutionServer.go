package main

import (
	"encoding/json"
	"fmt"
	b "github.com/martinre2/Evolution/Board"
	s "github.com/martinre2/Evolution/Stage"
	"log"
	"net/http"
	"strconv"
)

type Points struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}
type ExportConf struct {
	Fitness   int       `json:"fitness"`
	Radiogiro float64   `json:"radiogiro"`
	Dmax      float64   `json:"dmax"`
	Proteins  []*Points `json:"proteins"`
}

type Charts struct {
	Sum        []int     `json:"sum"`
	Average    []float32 `json:"average"`
	Best       []int     `json:"best"`
	Worts      []int     `json:"worts"`
	Gene       []int     `json:"gene"`
	RadioGiroP []float64 `json:"radioGiroP"`
	DmaxP      []float64 `json:"dMaxP"`
}

type ExportResponse struct {
	ConfPlot    []*ExportConf  `json:"confplot"`
	Generations []b.Generation `json:"generations"`
	Chart       Charts         `json:"charts"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println(r.FormValue("hpseq"))

		l, _ := strconv.Atoi(r.FormValue("lattice"))
		c, _ := strconv.Atoi(r.FormValue("numberconf"))
		run, _ := strconv.Atoi(r.FormValue("run"))
		pcx, _ := strconv.ParseFloat(r.FormValue("crossprob"), 32)
		experiments, _ := strconv.Atoi(r.FormValue("manyexp"))

		fmt.Println("Params", l, c, run)

		var bestBoard *b.BlackBoard
		var board *b.BlackBoard
		//var ex_randomStage s.RandomGenerationStage
		//var ex_simpleStage s.SimpleGenerationStage

		params := b.BlackBoardParameters{
			LatticeType:      l,
			NumConformations: c,
			NumGenerations:   run,
			NumExperiments:   1,
			OtherTechs:       []int{1, 0, 0},
			CrossOverProb:    float32(pcx),
			CrossOverOp:      1,
			MutationOp:       1,
		}

		bestBoard = b.NewBoard(params)

		for i := 0; i < experiments; i++ {
			board := b.NewBoard(params)
			board.Hplevel.HpString = r.FormValue("hpseq")

			ex_randomStage := new(s.RandomGenerationStage).Execute(board)
			ex_simpleStage := new(s.SimpleGenerationStage).Execute(board)

			_ = ex_randomStage
			_ = ex_simpleStage

			var bestFit int = 0

			for _, c := range board.Generations[len(board.Generations)-1].Conformations {

				if c.Fitness < bestFit {
					bestFit = c.Fitness
				}
			}

			board.BestFitness = bestFit

			if board.BestFitness < bestBoard.BestFitness {
				bestBoard = board
			}
		}

		for i := 0; i < experiments; i++ {
			fmt.Println(">>>i", i)
		}
		_ = bestBoard
		_ = board

		fmt.Println("<End Simulation")

		var export []*ExportConf
		var rs ExportResponse

		for _, c := range bestBoard.Generations[len(bestBoard.Generations)-1].Conformations {
			//for _, c := range board.Generations[0].Conformations {
			e := new(ExportConf)
			e.Fitness = c.Fitness
			e.Radiogiro = c.RadioGiroRCDM
			e.Dmax = c.Dmax
			for _, p := range c.Conformations {
				iP := new(Points)
				iP.X = p.XValue
				iP.Y = p.YValue
				iP.Z = p.ZValue
				e.Proteins = append(e.Proteins, iP)
			}
			export = append(export, e)
		}

		//Stuff for charts
		var ch Charts
		//Fitness Chart
		//Build Series
		sum := make([]int, len(bestBoard.Generations))
		average := make([]float32, len(bestBoard.Generations))
		best := make([]int, len(bestBoard.Generations))
		worts := make([]int, len(bestBoard.Generations))
		gene := make([]int, len(bestBoard.Generations))

		//Best
		for i, g := range bestBoard.Generations {
			gene[i] = i
			totalFit, totalBestFit, totalWortsFit := 0, 0, -100

			for _, c := range g.Conformations {
				totalFit += c.Fitness

				//totalBestFit = (totalBestFit > c.Fitness) ? (c.Fitness) : (totalBestFit)
				if totalBestFit > c.Fitness {
					totalBestFit = c.Fitness
				} else {
					totalBestFit = totalBestFit
				}
				//totalWortsFit = (totalWortsFit < c.Fitness) ? (c.Fitness) : (totalWortsFit)
				if totalWortsFit < c.Fitness {
					totalWortsFit = c.Fitness
				} else {
					totalWortsFit = totalWortsFit
				}
			}
			sum[i] = totalFit
			average[i] = float32(totalFit) / float32(len(g.Conformations))
			best[i] = totalBestFit
			worts[i] = totalWortsFit

		}

		//Radio Chart
		//Build Series
		radgiroP := make([]float64, len(bestBoard.Generations))
		dmaxP := make([]float64, len(bestBoard.Generations))

		for i, g := range bestBoard.Generations {
			radgiroP[i] = g.RadioGiroP
			dmaxP[i] = g.DmaxP
		}
		//end radio chart

		ch.Sum = sum
		ch.Average = average
		ch.Best = best
		ch.Worts = worts
		ch.Gene = gene
		ch.RadioGiroP = radgiroP
		ch.DmaxP = dmaxP

		fmt.Println(export)

		rs.ConfPlot = export
		//rs.Generations = bestBoard.Generations
		rs.Chart = ch

		//js, err := json.Marshal(export)
		js, err := json.Marshal(rs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		//fmt.Println("___>Peticion")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
