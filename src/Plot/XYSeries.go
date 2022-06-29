package GoPlot

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func CreateTimeSeriesPlot(dataset [][]float64, x_axis string, y_axis string, chart_title string, file_title string) {

	// Create Object
	p := plot.New()
	p.Add(plotter.NewGrid())

	// Set Parameters
	p.Title.Text = chart_title
	p.X.Label.Text = x_axis
	p.Y.Label.Text = y_axis

	// Fill Dataset
	for i := 0; i < len(dataset); i++ {

		pts := make(plotter.XYs, len(dataset[i]))

		for j := 0; j < len(dataset[i]); j++ {

			pts[j].X = float64(j)
			pts[j].Y = dataset[i][j]

		}

		err := plotutil.AddLinePoints(p, "", pts)

		if err != nil {
			log.Println("Data Cannot Be Added...")
		}

	}

	// Create File
	err := p.Save(4*vg.Inch, 4*vg.Inch, (file_title + ".png"))

	if err != nil {
		log.Println("Error Creating Time Series Chart")
	}

}
