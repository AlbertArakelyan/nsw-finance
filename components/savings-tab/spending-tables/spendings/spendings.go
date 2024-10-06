package spendings

import (
	"log"
	"nsw-finance/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type UIComponents struct {
	SpendingsList *fyne.Container
}

type Spendings struct {
	DB           repository.Repository
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	UIComponents UIComponents
}

func (spendings *Spendings) GetSpendingsContainer(spendingTableId int64) *fyne.Container {
	spendingsList := spendings.getSpenidingsList(spendingTableId)
	spendingsContainer := container.NewVBox(
		spendingsList,
		spendings.getAddSpendingButton(spendingTableId, spendingsList),
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
	spendings.UIComponents.SpendingsList = spendingsList

	spendings.appendSpendingToList(spendingsList, spendingsSlice)

	return spendingsList
}

func (spendings *Spendings) getAddSpendingButton(spendingTableId int64, spendingsList *fyne.Container) *widget.Button {
	addSpendingButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		err := spendings.AddNewSpending(spendingTableId, spendingsList)
		if err != nil {
			spendings.ErrorLog.Println(err)
			// log.Panic(err)
		}
	})

	return addSpendingButton
}
