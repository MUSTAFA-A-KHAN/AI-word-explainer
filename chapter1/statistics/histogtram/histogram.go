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
	// Create a new dataframe from a CSV file
	irisFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	irisDf := dataframe.ReadCSV(irisFile)

	defer irisFile.Close()
	for _, colName := range irisDf.Names() {
		if colName != "species" {
			v := make(plotter.Values, irisDf.Nrow())
			for i, floatVal := range irisDf.Col(colName).Float() {
				v[i] = floatVal
			}
			//make a plot and set its title
			p := plot.New()
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}

			h.Normalize(1)

			p.Add(h)
			filename := colName + ".png"
			if err := p.Save(4*vg.Inch, 4*vg.Inch, filename); err != nil {
				log.Fatal(err)
			}

		}

	}
}
