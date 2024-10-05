package spendingtables

import (
	"log"
	"nsw-finance/components/savings-tab/spending-tables/spendings"
	"nsw-finance/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SpendingTablesChildren struct {
	Spendings *spendings.Spendings
}

type UIComponents struct {
	AddNewSpendingTableEntryContainer *fyne.Container
	SpendingTablesContent             *fyne.Container
}

type SpendingTables struct {
	DB                         repository.Repository
	InfoLog                    *log.Logger
	ErrorLog                   *log.Logger
	UIComponents               UIComponents
	Children                   SpendingTablesChildren
	SpendingTablesSlice        []repository.SQLiteRepository
	IsSpendingTablesSliceEmpty bool
}

func (spendingTables *SpendingTables) GetSpendingTablesContainer() *fyne.Container {
	spendings := &spendings.Spendings{
		DB:       spendingTables.DB,
		InfoLog:  spendingTables.InfoLog,
		ErrorLog: spendingTables.ErrorLog,
	}
	spendingTables.Children.Spendings = spendings

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

	spendingTablesContainer := container.NewStack(
		container.NewVBox(
			spendingTablesHeader,
			widget.NewSeparator(), // TODO replace with primary color (rectangle from app-header.go)
		),
		container.NewVBox(addNewSpendingTableEntryContainer),
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
	addNewSpendingTableEntryContainer.MinSize().Subtract(fyne.NewSize(0, 10))

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
		spnedingTableLabel := canvas.NewText(spendingTable.Label, nil)
		spnedingTableLabel.TextStyle = fyne.TextStyle{Bold: true}
		c := container.NewVBox(
			container.NewHBox(
				spnedingTableLabel,
				// TODO Add Delete button
			),
			spendingTables.Children.Spendings.GetSpendingsContainer(spendingTable.ID), // TODO add this logic to the place when the list(table) is updated
			widget.NewSeparator(),
		)
		
		spendingTablesContent.Add(c)
	}

	spendingTables.IsSpendingTablesSliceEmpty = len(spendingTablesSlice) == 0

	spendingTablesScroll := container.NewVScroll(spendingTablesContent)
	spendingTablesScroll.SetMinSize(fyne.Size{Height: 300})
	spendingTablesBorderedContent := container.NewBorder(nil, nil, nil, nil, spendingTablesScroll)

	return spendingTablesBorderedContent
}
