package main

import (
	"log"
	"strconv"
)

func (app *App) GetSavingAmounts() (int64, int64, error) {
	saving, err := app.DB.GetSaving()
	if err != nil {
		app.Utils.ErrorLog.Println(err)
		log.Panic(err)
	}

	return saving.Amount, saving.AvailableAmount, nil
}

func (app *App) UpdateSavingAmount(amountText string) error {
	amount, err := strconv.ParseFloat(amountText, 64)
	if err != nil {
		app.Utils.ErrorLog.Println(err)
		log.Panic(err)
	}
	err = app.DB.UpdateSavingAmount(int64(amount))
	if err != nil {
		return err
	}

	return nil
}
