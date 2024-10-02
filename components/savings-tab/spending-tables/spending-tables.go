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

type UIComponents struct {
	AddNewSpendingTableEntryContainer *fyne.Container
}

type SpendingTables struct {
	DB           repository.Repository
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	UIComponents UIComponents
}

func (spendingTables *SpendingTables) GetSpendingTablesContainer() *fyne.Container {
	spendingTablesContainer := container.NewVBox(
		spendingTables.getSpendingTablesHeader(),
	)

	return spendingTablesContainer
}

func (spendingTables *SpendingTables) getSpendingTablesHeader() *fyne.Container {
	addButton := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
		spendingTables.UIComponents.AddNewSpendingTableEntryContainer.Show()
	})
	spendingTablesHeader := container.NewHBox(
		canvas.NewText("Spending Tables", nil),
		addButton,
	)
	addNewSpendingTableEntryContainer := spendingTables.getAddNewSpendingTableEntryContainer()
	addNewSpendingTableEntryContainer.Hide()

	spendingTablesContainer := container.NewVBox(
		spendingTablesHeader,
		widget.NewSeparator(),
		addNewSpendingTableEntryContainer,
	)

	return spendingTablesContainer
}

func (spendingTables *SpendingTables) getAddNewSpendingTableEntryContainer() *fyne.Container {
	addNewSpendingTableEntry := widget.NewEntry()
	addNewSpendingTableEntry.PlaceHolder = "Spending Table Name"

	saveButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {})

	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		spendingTables.UIComponents.AddNewSpendingTableEntryContainer.Hide()
	})
	deleteButton.Importance = widget.DangerImportance

	addNewSpendingTableEntryContainer := container.NewBorder(
		nil,
		nil,
		nil,
		container.NewHBox(saveButton, deleteButton),
		addNewSpendingTableEntry,
	)

	spendingTables.UIComponents.AddNewSpendingTableEntryContainer = addNewSpendingTableEntryContainer

	return addNewSpendingTableEntryContainer
}
