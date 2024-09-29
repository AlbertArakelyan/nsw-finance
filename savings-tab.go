package main

import (
	"fmt"
	"log"
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

	app.uiComponents.SavingsContainer = savingsContainer

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

	amountEntryContainer := container.NewBorder(nil, nil, amountText, nil, amountEntry)

	saveBtn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		err := app.UpdateSavingAmount(amountEntry.Text)
		if err != nil {
			app.utils.ErrorLog.Println(err)
			log.Panic(err)
			return
		}

		amountEntry.Refresh()

		// TODO also add logic for updating spendings (in tables)
	})
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

	return savingsContainer
}
