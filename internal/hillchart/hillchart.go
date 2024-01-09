package hillchart

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func GenerateHillChart() error {
	p := plot.New()
	p.Title.Text = "Hill Chart"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	xMin := -3.0
	xMax := 3.0
	pts := make(plotter.XYs, 100)
	for i := range pts {
		// x := xMin + (xMax-xMin)*float64(i)*2*math.Pi/100
		x := xMin + (xMax-xMin)*float64(i)/99
		pts[i].X = x
		pts[i].Y = 1 / (math.Sqrt(2 * math.Pi)) * math.Exp(-math.Pow(x, 2)/2)
	}

	// Create a line plotter with points
	line, err := plotter.NewLine(pts)
	if err != nil {

		panic(err)
	}
	p.Add(line)

	// Save the plot to a PNG file
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "normal.png"); err != nil {
		panic(err)
	}

	return nil
}
