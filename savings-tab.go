package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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
	amountEntry.Validator = func(s string) error {
		_, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		return nil
	}
	amountEntry.OnChanged = func(s string) {
		_, err := strconv.ParseFloat(s, 64)
		if err != nil {
			amountEntry.Text = amount
			return
		}
	}

	amountEntryContainer := container.NewBorder(nil, nil, amountText, nil, amountEntry)

	saveBtn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {})
	saveBtn.Alignment = widget.ButtonAlignTrailing
	saveBtn.Importance = widget.HighImportance

	savingsContainer := container.NewVBox(
		amountEntryContainer,
		container.NewHBox(
			availableAmount,
			layout.NewSpacer(),
			saveBtn,
		),
	)
	app.uiComponents.SavingsContainer = savingsContainer

	return savingsContainer
}
