package spendings

import (
	"log"
	"nsw-finance/repository"
	"strconv"

	"fyne.io/fyne/v2"
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
		spendingLabelEntry := widget.NewEntry()
		spendingLabelEntry.SetText(spending.Label)
		spendingLabelEntry.OnChanged = func(s string) {
			if s == "" {
				s = "New Spending"
				spendingLabelEntry.SetText(s)
			}
			err := spendings.UpdateSpendingLabel(spending.ID, s)
			if err != nil {
				spendings.ErrorLog.Println(err)
				// log.Panic(err)
			}
		}

		amountText := strconv.FormatFloat(float64(spending.Amount), 'f', 2, 64)
		spendingAmountEntry := widget.NewEntry()
		spendingAmountEntry.SetText(amountText)
		spendingAmountEntry.Validator = spendings.spendingAmountValidator
		spendingAmountEntry.OnChanged = func(s string) {
			err := spendings.ValidateAndUpdateSpendingAmount(spending.ID, s)
			if err != nil {
				spendings.ErrorLog.Println(err)
				// log.Panic(err)
			}
		}

		c := container.NewGridWithColumns(
			2,
			spendingLabelEntry,
			spendingAmountEntry,
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
