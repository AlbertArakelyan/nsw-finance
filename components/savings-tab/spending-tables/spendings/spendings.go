package spendings

import (
	"log"
	"nsw-finance/repository"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type UIComponents struct{}

type Spendings struct {
	DB           repository.Repository
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	UIComponents UIComponents
}

func (spendings *Spendings) GetSpendingsContainer(spendingTableId int64) *fyne.Container {
	spendingsContainer := container.NewVBox(
		spendings.getSpenidingsList(spendingTableId),
		spendings.getAddSpendingButton(spendingTableId),
	)

	return spendingsContainer
}

func (spendings *Spendings) getSpenidingsList(savingTableId int64) *fyne.Container {
	spendingsSlice, err := spendings.DB.GetSpendings(savingTableId)
	if err != nil {
		spendings.ErrorLog.Println(err)
		// log.Panic(err)
	}

	spendingsList := container.NewVBox()
	for _, spending := range spendingsSlice {
		amountText := strconv.FormatFloat(float64(spending.Amount), 'f', 2, 64)

		c := container.NewHBox(
			canvas.NewText(spending.Label, nil), // TODO change into entry and update on DB on change
			canvas.NewText(amountText, nil), // TODO change into entry and update on DB on change
		)
		spendingsList.Add(c)
	}

	return spendingsList
}

func (spendings *Spendings) getAddSpendingButton(spendingTableId int64) *widget.Button {
	addSpendingButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		err := spendings.AddNewSpending(spendingTableId)
		if err != nil {
			spendings.ErrorLog.Println(err)
			// log.Panic(err)
		}
	})

	return addSpendingButton
}
