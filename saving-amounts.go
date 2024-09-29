package main

import "log"

func (app *App) GetSavingAmounts() (int64, int64, error) {
	saving, err := app.DB.GetSaving()
	if err != nil {
		app.utils.ErrorLog.Println(err)
		log.Panic(err)
	}

	return saving.Amount, saving.AvailableAmount, nil
}
