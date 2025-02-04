package main

import (
	"image/color"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// A box plot is a special type of diagram that shows the quartiles in a box and the line extending from the lowest to the highest value.
func main() {
	irisFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)

	p := plot.New()

	p.Title.Text = "Iris Dataset"
	p.BackgroundColor = color.Transparent
	p.Y.Label.Text = "Values"
	p.Title.TextStyle.Color = color.White

	// Create the box for our data
	w := vg.Points(50)

	for idx, colName := range irisDF.Names() {
		if colName != "species" {
			v := make(plotter.Values, irisDF.Nrow())

			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			b, err := plotter.NewBoxPlot(w, float64(idx), v)
			if err != nil {
				log.Fatal(err)
			}
			p.Add(b)

		}

	}

	p.NominalX("sepal_length", "sepal_width", "petal_length", "petal_width")
	if err := p.Save(6*vg.Inch, 8*vg.Inch, "boxplots.png"); err != nil {
		log.Fatal(err)
	}
}
