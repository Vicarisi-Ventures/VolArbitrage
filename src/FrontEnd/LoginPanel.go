package FrontEnd

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoginPanel(c1 chan *fyne.Container, c2 chan []string) {

	var api_key string
	var account_id string

	// Enter Api Key
	api_label := widget.NewLabel("Enter Api Key: ")
	api_input := widget.NewEntry()
	api_input.SetPlaceHolder("Enter text...")

	// Enter Account ID
	account_label := widget.NewLabel("Enter Account ID: ")
	account_input := widget.NewEntry()
	account_input.SetPlaceHolder("Enter text...")

	initial_content := container.NewVBox(

		api_label,

		api_input, widget.NewButton("Enter",
			func() {

				api_key = api_input.Text
				log.Println("Api Key:", api_input.Text)

			}),

		account_label,

		account_input, widget.NewButton("Enter",
			func() {

				account_id = account_input.Text
				log.Println("Account ID:", account_input.Text)

			}),
	)

	c1 <- initial_content
	c2 <- []string{api_key, account_id}

}
