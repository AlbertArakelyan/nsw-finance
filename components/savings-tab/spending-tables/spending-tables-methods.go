package spendingtables

import (
	"errors"
	"nsw-finance/repository"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (spendingTables *SpendingTables) ValidateAndAddSpendingTable(label string, savingId int64) error {
	if label == "" {
		return errors.New("label cannot be empty")
	}

	lastInsertId, err := spendingTables.DB.AddSpendingTable(label, savingId)
	if err != nil {
		return err
	}

	spendingTables.RefreshSpendingsTablesContent(lastInsertId)

	return nil
}

func (spendingTables *SpendingTables) GetSpendingTables(savingId int64) ([]repository.SpendingTable, error) {
	spendingTablesSlice, err := spendingTables.DB.GetSpendingTables(savingId)
	if err != nil {
		return nil, err
	}

	if len(spendingTablesSlice) == 0 {
		return nil, errors.New("no spending tables found")
	}

	return spendingTablesSlice, nil
}

func (spendingTables *SpendingTables) RefreshSpendingsTablesContent(lastSpendingTableId int64) {
	lastSpendingTable, err := spendingTables.DB.GetSpendingTableByID(lastSpendingTableId)
	if err != nil {
		spendingTables.ErrorLog.Println(err)
		return
	}

	c := container.NewVBox(
		container.NewHBox(
			canvas.NewText(lastSpendingTable.Label, nil),
			// TODO Add Delete button
		),
		spendingTables.Children.Spendings.GetSpendingsContainer(),
		widget.NewSeparator(),
	)

	if spendingTables.IsSpendingTablesSliceEmpty {
		spendingTables.UIComponents.SpendingTablesContent.RemoveAll()
		spendingTables.UIComponents.SpendingTablesContent.Add(c)
		spendingTables.IsSpendingTablesSliceEmpty = false
	} else {
		spendingTables.UIComponents.SpendingTablesContent.Add(c)
	}
}
