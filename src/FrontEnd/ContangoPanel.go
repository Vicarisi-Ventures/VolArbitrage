package FrontEnd

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ContangoPanel() *fyne.Container {

	contango_label := widget.NewLabel("Equity Risk Premium: ")

	risk_premia_label := widget.NewLabel("Risk Premia Backtest")

	input_label := widget.NewLabel("Enter Ticker Symbol: ")
	input := widget.NewEntry()
	input.SetPlaceHolder("Ex: AAPL")

	button := widget.NewButton("Enter", func() {
		ticker_symbol = input.Text
		log.Println("Ticker Symbol: ", ticker_symbol)
	})

	temp_left := container.NewVBox(

		input_label,
		input,
		button,
	)

	display_premia_grid := container.New((layout.NewGridLayout(2)), temp_left, risk_premia_label)

	r, _ := LoadResourceFromPath("C:/Users/18627/Documents/GitHub/VolArbitrage/VIXFutures.tqqq_test.png")
	image := canvas.NewImageFromResource(r)
	image.FillMode = canvas.ImageFillOriginal

	p := container.NewVBox(

		contango_label,
		display_premia_grid,
		image,
	)

	return p

}
