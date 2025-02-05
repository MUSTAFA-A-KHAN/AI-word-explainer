package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	adv, err := os.Open("../data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer adv.Close()
	advDF := dataframe.ReadCSV(adv)

	for _, colName := range advDF.Names() {
		plotVals := make(plotter.Values, advDF.Nrow())

		for i, floatVal := range advDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}
		p := plot.New() // Create a new plot
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}
		h.Normalize(1)
		p.Add(h)

		if err := p.Save(4*vg.Inch, 4*vg.Inch, "./output/"+colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}

}
