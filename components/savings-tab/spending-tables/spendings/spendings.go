package spendings

import (
	"log"
	"nsw-finance/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type UIComponents struct {
	SpendingEntryContainer *fyne.Container
}

type Spendings struct {
	DB           repository.Repository
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	UIComponents UIComponents
}

func (spendings *Spendings) GetSpendingsContainer() *fyne.Container {
	spendingsContainer := container.NewVBox(
		canvas.NewText("Spendings", nil),
		spendings.getAddSpendingButton(),
	)

	return spendingsContainer
}

func (spendings *Spendings) getAddSpendingButton() *widget.Button {
	addSpendingButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {})

	return addSpendingButton
}
