package main

import (
	"fmt"
	"log"
	"strconv"

	"fyne.io/fyne/v2/canvas"
)

func (app *App) getSavingsText() (string, *canvas.Text) {
	var availableAmount *canvas.Text

	// Get savings info from DB
	savingAmount, savingAvailableAmount, err := app.GetSavingAmounts()
	if err != nil {
		app.utils.ErrorLog.Println(err)
		log.Panic(err)
	}

	amountText := strconv.FormatFloat(float64(savingAmount), 'f', 2, 64)
	availableAmountText := fmt.Sprintf("Available: %.2f %s", float64(savingAvailableAmount), currency)

	availableAmount = canvas.NewText(availableAmountText, appTheme.colors.gray)

	return amountText, availableAmount
}
