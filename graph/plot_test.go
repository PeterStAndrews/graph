package graph

import (
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func PercolateExperiment(n int) plotter.XYs {

	var Ts []float64 = linspace(0.0, 1.0, n)

	var p Poisson    // poisson degree distribution
	p.kmean = 6      // average degree
	const N = 100000 // number of ndoes

	ns := p.generateSamples(N)
	pts := make(plotter.XYs, n)
	for i := range pts {
		pts[i].X = Ts[i]
		pts[i].Y = BondPercolation(ConfigurationModel(&ns), Ts[i])
	}
	return pts
}

func TestBondPercolationPlot(t *testing.T) {
	p := plot.New()

	p.Title.Text = "Percolation experiment"
	p.X.Label.Text = "T"
	p.Y.Label.Text = "S"

	err := plotutil.AddLinePoints(p,
		"First", PercolateExperiment(25))

	if err != nil {
		panic(err)
	}

	// Save the plot to a PDF file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "percolation.pdf"); err != nil {
		panic(err)
	}

}
