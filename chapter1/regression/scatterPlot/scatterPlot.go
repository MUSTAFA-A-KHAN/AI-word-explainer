package main

import (
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
	advDF := dataframe.ReadCSV(adv)

	//Extract the target columns(Sales)
	yVals := advDF.Col("Sales").Float()

	for _, colName := range advDF.Names() {
		pts := make(plotter.XYs, advDF.Nrow())

		//fill pts with Data
		for i, floatVal := range advDF.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}
		p := plot.New() //create a new plot
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())
		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Radius = vg.Points(3)
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, "./output/"+colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}
}
