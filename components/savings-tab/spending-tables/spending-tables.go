package spendingtables

import (
	"log"
	"nsw-finance/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type SpendingTables struct {
	DB       repository.Repository
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (spendingTables *SpendingTables) GetSpendingTablesContainer() *fyne.Container {
	return container.NewVBox(canvas.NewText("Spending Tables Container", nil))
}
