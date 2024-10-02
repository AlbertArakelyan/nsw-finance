package savingstab

import (
	"fmt"
	"log"
	"nsw-finance/theme"
	"strconv"

	"fyne.io/fyne/v2/canvas"
)

func (savingsTab *SavingsTab) getSavingsText() (string, *canvas.Text) {
	var availableAmount *canvas.Text

	// Get savings info from DB
	savingAmount, savingAvailableAmount, err := savingsTab.GetSavingAmounts()
	if err != nil {
		savingsTab.ErrorLog.Println(err)
		log.Panic(err)
	}

	amountText := strconv.FormatFloat(float64(savingAmount), 'f', 2, 64)
	availableAmountText := fmt.Sprintf("Available: %.2f %s", float64(savingAvailableAmount), currency)

	availableAmount = canvas.NewText(availableAmountText, theme.AppTheme.Colors.Gray)

	return amountText, availableAmount
}
