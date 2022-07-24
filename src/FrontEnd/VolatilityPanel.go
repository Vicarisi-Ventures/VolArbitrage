package FrontEnd

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var ticker_symbol string

func ImpliedVolatilityPanel() *fyne.Container {

	volatility_label := widget.NewLabel("Implied Volatility Curve")

	input_label := widget.NewLabel("Enter Ticker Symbol: ")
	input := widget.NewEntry()
	input.SetPlaceHolder("Ex: AAPL")

	button := widget.NewButton("Enter", func() {
		ticker_symbol = input.Text
		log.Println("Ticker Symbol: ", ticker_symbol)
	})

	var check []*widget.Check

	check = append(check, widget.NewCheck("Optional", func(value bool) {
		log.Println("Check set to", value)
	}))

	left := container.NewVBox(

		input_label,
		input,
		button,
	)

	for i := 0; i < len(check); i++ {
		left.Add(check[i])
	}

	right := container.NewVBox(

		volatility_label,
	)

	grid := container.New(layout.NewGridLayout(2), left, right)

	p := container.NewVBox(

		grid,
	)

	return p

}
