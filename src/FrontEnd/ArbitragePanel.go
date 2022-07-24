package FrontEnd

import (
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	m "v2/src/Mongo"
)

func ScreenerPanel() *fyne.Container {

	// Initialize Client
	client := m.GetMongoConnection()
	ticker_symbols := m.GetStockTickers(true)

	// Realized Volatility
	var hv30 [][]float64
	var hv60 [][]float64
	var hv90 [][]float64
	var hv120 [][]float64

	// Variance Risk Premium
	var vrp30 []float64
	var vrp60 []float64
	var vrp90 []float64
	var vrp120 []float64

	// Iterate Over Data
	for i := 0; i < len(ticker_symbols); i++ {

		Data := m.FetchVolatilityMongoDB(client, ticker_symbols[i])

		var temp30 []float64
		for j := 0; j < len(Data.HV.HV30); j++ {
			temp30 = append(temp30, Data.HV.HV30[j])
		}
		hv30 = append(hv30, temp30)

		var temp60 []float64
		for j := 0; j < len(Data.HV.HV60); j++ {
			temp60 = append(temp60, Data.HV.HV60[j])
		}
		hv60 = append(hv60, temp60)

		var temp90 []float64
		for j := 0; j < len(Data.HV.HV90); j++ {
			temp90 = append(temp90, Data.HV.HV90[j])
		}
		hv90 = append(hv90, temp90)

		var temp120 []float64
		for j := 0; j < len(Data.HV.HV120); j++ {
			temp120 = append(temp120, Data.HV.HV120[j])
		}
		hv120 = append(hv120, temp120)

		vrp30 = append(vrp30, Data.VRP.VRP30)
		vrp60 = append(vrp60, Data.VRP.VRP60)
		vrp90 = append(vrp90, Data.VRP.VRP90)
		vrp120 = append(vrp120, Data.VRP.VRP120)

	}

	// Initialize Variables
	var x [][]string
	var top_tickers [][]string
	top_index30, top_five_vrp30 := findTopFive(vrp30)
	top_index60, top_five_vrp60 := findTopFive(vrp60)
	top_index90, top_five_vrp90 := findTopFive(vrp90)
	top_index120, top_five_vrp120 := findTopFive(vrp120)

	for i := 0; i < 5; i++ {

		var temp_x []string
		var temp_t []string

		// 30
		temp_t = append(temp_t, ticker_symbols[top_index30[i]]+": ")
		temp_x = append(temp_x, strconv.FormatFloat(top_five_vrp30[i], 'f', 3, 64))

		// 60
		temp_t = append(temp_t, ticker_symbols[top_index60[i]]+": ")
		temp_x = append(temp_x, strconv.FormatFloat(top_five_vrp60[i], 'f', 3, 64))

		// 90
		temp_t = append(temp_t, ticker_symbols[top_index90[i]]+": ")
		temp_x = append(temp_x, strconv.FormatFloat(top_five_vrp90[i], 'f', 3, 64))

		// 120
		temp_t = append(temp_t, ticker_symbols[top_index120[i]]+": ")
		temp_x = append(temp_x, strconv.FormatFloat(top_five_vrp120[i], 'f', 3, 64))

		x = append(x, temp_x)
		top_tickers = append(top_tickers, temp_t)

	}

	// Initialize Variables
	var y [][]string
	var bottom_tickers [][]string
	bottom_index30, bottom_five_vrp30 := findBottomFive(vrp30)
	bottom_index60, bottom_five_vrp60 := findBottomFive(vrp60)
	bottom_index90, bottom_five_vrp90 := findBottomFive(vrp90)
	bottom_index120, bottom_five_vrp120 := findBottomFive(vrp120)

	for i := 0; i < 5; i++ {

		var temp_y []string
		var temp_t []string

		// 30
		temp_t = append(temp_t, ticker_symbols[bottom_index30[i]]+": ")
		temp_y = append(temp_y, strconv.FormatFloat(bottom_five_vrp30[i], 'f', 3, 64))

		// 60
		temp_t = append(temp_t, ticker_symbols[bottom_index60[i]]+": ")
		temp_y = append(temp_y, strconv.FormatFloat(bottom_five_vrp60[i], 'f', 3, 64))

		// 90
		temp_t = append(temp_t, ticker_symbols[bottom_index90[i]]+": ")
		temp_y = append(temp_y, strconv.FormatFloat(bottom_five_vrp90[i], 'f', 3, 64))

		// 120
		temp_t = append(temp_t, ticker_symbols[bottom_index120[i]]+": ")
		temp_y = append(temp_y, strconv.FormatFloat(bottom_five_vrp120[i], 'f', 3, 64))

		y = append(y, temp_y)
		bottom_tickers = append(bottom_tickers, temp_t)

	}

	screener_top_vrp_label := widget.NewLabel("Top Five Highest Variance Risk Premium: ")
	screener_bottom_vrp_label := widget.NewLabel("Top Five Lowest Variance Risk Premium: ")
	title_grid := container.New(layout.NewGridLayout(2), screener_top_vrp_label, screener_bottom_vrp_label)

	a := container.NewVBox(
		widget.NewLabel("30 Day: "),
		widget.NewLabel(top_tickers[0][0]+x[0][0]),
		widget.NewLabel(top_tickers[1][0]+x[1][0]),
		widget.NewLabel(top_tickers[2][0]+x[2][0]),
		widget.NewLabel(top_tickers[3][0]+x[3][0]),
		widget.NewLabel(top_tickers[4][0]+x[4][0]),
	)

	b := container.NewVBox(
		widget.NewLabel("60 Day: "),
		widget.NewLabel(top_tickers[0][1]+x[0][1]),
		widget.NewLabel(top_tickers[1][1]+x[1][1]),
		widget.NewLabel(top_tickers[2][1]+x[2][1]),
		widget.NewLabel(top_tickers[3][1]+x[3][1]),
		widget.NewLabel(top_tickers[4][1]+x[4][1]),
	)

	c := container.NewVBox(
		widget.NewLabel("90 Day: "),
		widget.NewLabel(top_tickers[0][2]+x[0][2]),
		widget.NewLabel(top_tickers[1][2]+x[1][2]),
		widget.NewLabel(top_tickers[2][2]+x[2][2]),
		widget.NewLabel(top_tickers[3][2]+x[3][2]),
		widget.NewLabel(top_tickers[4][2]+x[4][2]),
	)

	d := container.NewVBox(
		widget.NewLabel("120 Day: "),
		widget.NewLabel(top_tickers[0][3]+x[0][3]),
		widget.NewLabel(top_tickers[1][3]+x[1][3]),
		widget.NewLabel(top_tickers[2][3]+x[2][3]),
		widget.NewLabel(top_tickers[3][3]+x[3][3]),
		widget.NewLabel(top_tickers[4][3]+x[4][3]),
	)

	left_data_grid := container.New(layout.NewGridLayout(4), a, b, c, d)

	alpha := container.NewVBox(
		widget.NewLabel("30 Day: "),
		widget.NewLabel(bottom_tickers[0][0]+y[0][0]),
		widget.NewLabel(bottom_tickers[1][0]+y[1][0]),
		widget.NewLabel(bottom_tickers[2][0]+y[2][0]),
		widget.NewLabel(bottom_tickers[3][0]+y[3][0]),
		widget.NewLabel(bottom_tickers[4][0]+y[4][0]),
	)

	beta := container.NewVBox(
		widget.NewLabel("60 Day: "),
		widget.NewLabel(bottom_tickers[0][1]+y[0][1]),
		widget.NewLabel(bottom_tickers[1][1]+y[1][1]),
		widget.NewLabel(bottom_tickers[2][1]+y[2][1]),
		widget.NewLabel(bottom_tickers[3][1]+y[3][1]),
		widget.NewLabel(bottom_tickers[4][1]+y[4][1]),
	)

	gamma := container.NewVBox(
		widget.NewLabel("90 Day: "),
		widget.NewLabel(bottom_tickers[0][2]+y[0][2]),
		widget.NewLabel(bottom_tickers[1][2]+y[1][2]),
		widget.NewLabel(bottom_tickers[2][2]+y[2][2]),
		widget.NewLabel(bottom_tickers[3][2]+y[3][2]),
		widget.NewLabel(bottom_tickers[4][2]+y[4][1]),
	)

	delta := container.NewVBox(
		widget.NewLabel("120 Day: "),
		widget.NewLabel(bottom_tickers[0][3]+y[0][3]),
		widget.NewLabel(bottom_tickers[1][3]+y[1][3]),
		widget.NewLabel(bottom_tickers[2][3]+y[2][3]),
		widget.NewLabel(bottom_tickers[3][3]+y[3][3]),
		widget.NewLabel(bottom_tickers[4][3]+y[4][3]),
	)

	right_data_grid := container.New(layout.NewGridLayout(4), alpha, beta, gamma, delta)
	data_grid := container.New(layout.NewGridLayout(2), left_data_grid, right_data_grid)

	left_chart_grid := container.NewVBox()
	right_chart_grid := container.NewVBox()
	risk_premia_label := widget.NewLabel("Historical Risk Premia")
	chart_grid := container.New(layout.NewGridLayout(2), left_chart_grid, right_chart_grid)

	chart_final := container.NewVBox(

		risk_premia_label,
		chart_grid,
	)

	input_label := widget.NewLabel("Enter Ticker Symbol: ")
	input := widget.NewEntry()
	input.SetPlaceHolder("Ex: AAPL")

	button := widget.NewButton("Enter", func() {

		ticker_symbol = input.Text
		log.Println("Ticker Symbol: ", ticker_symbol)

		if len(ticker_symbol) > 0 {

			left_chart_grid.RemoveAll()
			right_chart_grid.RemoveAll()

		}

	})

	temp_left := container.NewVBox(

		input_label,
		input,
		button,
	)

	display_premia_grid := container.New((layout.NewGridLayout(2)), temp_left, chart_final)

	p := container.NewVBox(
		title_grid,
		canvas.NewLine(color.Black),
		data_grid,
		canvas.NewLine(color.Black),
		display_premia_grid,
	)

	return p

}
