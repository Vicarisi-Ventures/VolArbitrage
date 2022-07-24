package FrontEnd

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SimulationPanel() *fyne.Container {

	simulation_label := widget.NewLabel("Monte Carlo Engine")

	p := container.NewVBox(

		simulation_label,
	)

	return p

}
