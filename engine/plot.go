package engine

import (
	"fmt"
	"github.com/mumax/3/graph"
	"log"
	"net/http"
)

func (g *guistate) servePlot(w http.ResponseWriter, r *http.Request) {

	log.Println("plotting...")

	// atomically get local copies to avoid race conditions
	var (
		data [][][]float64
	)
	InjectAndWait(func() {
		hist := Table.history
		data = make([][][]float64, len(hist))
		for i := range data {
			data[i] = make([][]float64, len(hist[i]))
			copy(data[i], hist[i]) // only copies slice headers, sufficient to avoid races (append to slices only)
		}
		if len(hist) > 0 {
			log.Println("len(table.hisotry[0]) = ", len(hist[0]))
			log.Println("len(data[0]) = ", len(data[0]))
		}
	})

	w.Header().Set("Content-Type", "image/svg+xml")
	plot := graph.New(w, 600, 300) // (!) canvas size duplicated in html.go
	defer plot.End()

	if len(data) == 0 || len(data[0]) == 0 || len(data[0][0]) < 2 { // need 2 points (from time column)
		//log.Println("no data to plot, len(data) = ", len(data))
		//	if len(data) > 0{
		//	log.Println("len(data[0]) = ", len(data[0]))
		//	}
		return
	}
	tMax := data[0][0][len(data[0][0])-1]
	plot.SetRanges(0, tMax, -1, 1)
	plot.DrawAxes(tMax/5, 0.5)
	plot.DrawXLabel("t (s)")
	for c := 0; c < 3; c++ {
		plot.LineStyle = fmt.Sprintf(`style="fill:none;stroke-width:2;stroke:%v"`, [3]string{"red", "green", "blue"}[c])
		plot.Polyline(data[0][0], data[1][c])
	}
}
