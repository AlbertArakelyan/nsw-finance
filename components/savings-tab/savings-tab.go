package savingstab

import (
	"fmt"
	"log"
	spendingtables "nsw-finance/components/savings-tab/spending-tables"
	"nsw-finance/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SavingsChildren struct{}

type SavingsTab struct {
	DB       repository.Repository
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var currency = "AMD"

func (savingsTab *SavingsTab) GetSavingsTab() *fyne.Container {
	savingsTextContainer := savingsTab.getSavingsContainer()

	spendingTables := &spendingtables.SpendingTables{
		DB:       savingsTab.DB,
		InfoLog:  savingsTab.InfoLog,
		ErrorLog: savingsTab.ErrorLog,
	}
	spendingTablesContainer := spendingTables.GetSpendingTablesContainer()

	savingsContainer := container.NewVBox(savingsTextContainer, spendingTablesContainer)

	return savingsContainer
}

func (savingsTab *SavingsTab) getSavingsContainer() *fyne.Container {
	amount, availableAmount := savingsTab.getSavingsText()

	// Amount text and entry
	amountText := canvas.NewText("Savings: ", nil)
	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder(fmt.Sprintf("Amount (%s)", currency))
	amountEntry.Text = amount
	amountEntry.OnChanged = func(s string) {
		savingsTab.ValidateAndUpdateSavingAmount(s)
	}
	amountEntry.Validator = amountEntryValidator

	amountEntryContainer := container.NewBorder(nil, nil, amountText, nil, amountEntry)

	saveBtn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		savingsTab.ValidateAndUpdateSavingAmount(amountEntry.Text)
		amountEntry.Refresh()

		// TODO also add logic for updating spendings (in tables)
	})
	saveBtn.Alignment = widget.ButtonAlignTrailing
	saveBtn.Importance = widget.HighImportance

	savingsContainer := container.NewVBox(
		amountEntryContainer,
		container.NewHBox(
			availableAmount,
			layout.NewSpacer(),
			saveBtn,
		),
	)

	return savingsContainer
}

// TODO add functions or reusable function for displaying spending tables