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

func main() {

	adv, err := os.Open("../../data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer adv.Close()

	advDF := dataframe.ReadCSV(adv)
	yVals := advDF.Col("Sales").Float()

	//values for plotting
	pts := make(plotter.XYs, advDF.Nrow())

	//predicted values
	ptsPred := make(plotter.XYs, advDF.Nrow())

	for i, floatVal := range advDF.Col("TV").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = predictedVal(floatVal)
	}
	p := plot.New() //create a new plot
	p.X.Label.Text = "TV"
	p.Y.Label.Text = "Sales"
	p.Add(plotter.NewGrid())

	// Add scatter plot
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Radius = vg.Points(3)
	s.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	p.Add(s, l)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "../output/Advertise_scatter.png"); err != nil {
		log.Fatal(err)
	}

}

// prediction function
// predict sales based on TV advertising using formula predicted using linear regression
func predictedVal(TV float64) float64 {
	return 7.0688 + TV*0.0489
}
