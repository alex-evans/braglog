package hillchart

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func GenerateHillChart(pointPercentage float64, pointLabel string) error {
	p := plot.New()
	p.Title.Text = "Hill Chart"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	xMin := -3.0
	xMax := 3.0
	pts := make(plotter.XYs, 100)
	for i := range pts {
		x := xMin + (xMax-xMin)*float64(i)/99
		pts[i].X = x
		pts[i].Y = 1 / (math.Sqrt(2 * math.Pi)) * math.Exp(-math.Pow(x, 2)/2)
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		return fmt.Errorf("failed to create plot: %v", err)
	}
	p.Add(line)

	pointX := xMin + (xMax-xMin)*pointPercentage/100

	point, err := plotter.NewScatter(plotter.XYLabels{
		XYs:    plotter.XYs{{pointX, 1 / (math.Sqrt(2 * math.Pi)) * math.Exp(-math.Pow(pointX, 2)/2)}},
		Labels: []string{pointLabel},
	})
	if err != nil {
		return fmt.Errorf("failed to create point on hillchart: %v", err)
	}
	p.Add(point)

	if err := p.Save(10*vg.Inch, 4*vg.Inch, "normal.png"); err != nil {
		return fmt.Errorf("failed to save hillchart image: %v", err)
	}

	return nil
}
