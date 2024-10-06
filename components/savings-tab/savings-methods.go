package savingstab

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2/canvas"
)

func (savingsTab *SavingsTab) GetSavingAmounts() (float64, int64, error) {
	saving, err := savingsTab.DB.GetSaving()
	if err != nil {
		savingsTab.ErrorLog.Println(err)
		log.Panic(err)
	}

	return saving.Amount, saving.AvailableAmount, nil
}

func (savingsTab *SavingsTab) UpdateSavingAmount(amountText string) error {
	amount, err := strconv.ParseFloat(amountText, 64)
	if err != nil {
		savingsTab.ErrorLog.Println(err)
		log.Panic(err)
	}

	err = savingsTab.DB.UpdateSavingAmount(int64(amount))
	if err != nil {
		return err
	}

	return nil
}

func (savingsTab *SavingsTab) UpdateAvailableSavingAmount(availableAmount *canvas.Text) error {
	_, err := savingsTab.DB.UpdateAvailableSavingAmount()
	if err != nil {
		return err
	}

	_, newAvailableAmount := savingsTab.getSavingsText()
	availableAmount.Text = newAvailableAmount.Text
	availableAmount.Refresh()

	return nil
}

func amountEntryValidator(s string) error {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	return nil
}

func (savingsTab *SavingsTab) ValidateAndUpdateSavingAmount(amountText string) error {
	err := amountEntryValidator(amountText)
	if err != nil {
		return err
	}

	err = savingsTab.UpdateSavingAmount(amountText)
	if err != nil {
		return err
	}

	return nil
}
