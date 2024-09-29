package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var currency = "AMD"

type Savings struct {
	Amount          float64 `json:"amount"`
	AvailableAmount float64 `json:"available_amount"`
}

func (app *App) savingsTab() *fyne.Container {
	savingsTextContainer := app.getSavingsTextContainer()
	savingsContainer := container.NewVBox(savingsTextContainer)

	return savingsContainer
}

func (app *App) getSavingsTextContainer() *fyne.Container {
	amount, availableAmount := app.getSavingsText()

	amountText := canvas.NewText("Savings: ", nil)
	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder(fmt.Sprintf("Amount (%s)", currency))
	amountEntry.Text = amount

	amountEntryContainer := container.NewBorder(nil, nil, amountText, nil, amountEntry)

	savingsContainer := container.NewVBox(
		amountEntryContainer,
		availableAmount,
	)
	app.uiComponents.SavingsContainer = savingsContainer

	return savingsContainer
}

func (app *App) getSavingsText() (string, *canvas.Text) {
	var availableAmount *canvas.Text

	// Get savings info from DB

	amountText := strconv.FormatFloat(280000, 'f', 2, 64)
	availableAmountText := fmt.Sprintf("Available: %.2f %s", float64(200000), currency)

	availableAmount = canvas.NewText(availableAmountText, appTheme.colors.gray)

	return amountText, availableAmount
}
