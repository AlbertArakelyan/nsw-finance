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
	SpendingTablesContent             *fyne.Container
}

type SpendingTables struct {
	DB                         repository.Repository
	InfoLog                    *log.Logger
	ErrorLog                   *log.Logger
	UIComponents               UIComponents
	SpendingTablesSlice        []repository.SQLiteRepository
	IsSpendingTablesSliceEmpty bool
}

func (spendingTables *SpendingTables) GetSpendingTablesContainer() *fyne.Container {
	spendingTablesContainer := container.NewVBox(
		spendingTables.getSpendingTablesHeader(),
		spendingTables.getSpendingTables(),
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
		widget.NewSeparator(), // TODO replace with primary color (rectangle from app-header.go)
		addNewSpendingTableEntryContainer,
	)

	return spendingTablesContainer
}

func (spendingTables *SpendingTables) getAddNewSpendingTableEntryContainer() *fyne.Container {
	// Maybe in future change whole entry with Dialog
	addNewSpendingTableEntry := widget.NewEntry()
	addNewSpendingTableEntry.PlaceHolder = "Spending Table Name"

	saveButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		err := spendingTables.ValidateAndAddSpendingTable(addNewSpendingTableEntry.Text, 1)
		if err != nil {
			spendingTables.ErrorLog.Println(err)
			return
		}

		addNewSpendingTableEntry.SetText("")
		spendingTables.UIComponents.AddNewSpendingTableEntryContainer.Hide()
	})

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

func (spendingTables *SpendingTables) getSpendingTables() *fyne.Container {
	spendingTablesContent := container.NewVBox()
	spendingTables.UIComponents.SpendingTablesContent = spendingTablesContent

	spendingTablesSlice, err := spendingTables.GetSpendingTables(1) // TODO change 1 to actual id
	if err != nil {
		spendingTables.IsSpendingTablesSliceEmpty = true
		spendingTables.ErrorLog.Println(err)
		spendingTablesContent.Add(container.NewCenter(canvas.NewText("No Spending Tables", nil)))
		return spendingTablesContent
	}

	spendingTablesContent.RemoveAll()

	for _, spendingTable := range spendingTablesSlice {
		c := container.NewVBox(
			container.NewHBox(
				canvas.NewText(spendingTable.Label, nil),
				// TODO Add Delete button
			),
			// TODO replace with some getTableContent method (which will return the table by itself)
			widget.NewSeparator(),
		)
		spendingTablesContent.Add(c)
	}

	spendingTables.IsSpendingTablesSliceEmpty = len(spendingTablesSlice) == 0

	return spendingTablesContent
}
