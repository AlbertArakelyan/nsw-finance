package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var currency = "AMD"

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
