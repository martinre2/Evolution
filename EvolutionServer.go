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

		fmt.Println("Params", l, c, run)

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

		board := b.NewBoard(params)
		board.Hplevel.HpString = r.FormValue("hpseq")

		ex_randomStage := new(s.RandomGenerationStage).Execute(board)
		ex_simpleStage := new(s.SimpleGenerationStage).Execute(board)

		_ = board
		_ = ex_randomStage
		_ = ex_simpleStage

		fmt.Println("<End Simulation")

		var export []*ExportConf

		for _, c := range board.Generations[len(board.Generations)-1].Conformations {
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

		fmt.Println(export)

		js, err := json.Marshal(export)
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
