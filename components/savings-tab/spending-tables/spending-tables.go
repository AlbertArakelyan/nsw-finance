package spendingtables

import (
	"log"
	"nsw-finance/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SpendingTables struct {
	DB       repository.Repository
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (spendingTables *SpendingTables) GetSpendingTablesContainer() *fyne.Container {
	spendingTablesContainer := container.NewVBox(
		spendingTables.getSpendingTablesHeader(),
	)

	return spendingTablesContainer
}

func (spendingTables *SpendingTables) getSpendingTablesHeader() *fyne.Container {
	addButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {})
	spendingTablesHeader := container.NewHBox(
		canvas.NewText("Spending Tables", nil),
		addButton,
	)

	return spendingTablesHeader
}
