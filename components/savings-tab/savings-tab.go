package savingstab

import (
	"fmt"
	"log"
	spendingtables "nsw-finance/components/savings-tab/spending-tables"
	savingsrepository "nsw-finance/repository/savings-repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SavingsChildren struct {
	SpendingTablesContainer *fyne.Container
}

type SavingsTab struct {
	DB       savingsrepository.Repository
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Children SavingsChildren
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
	savingsTab.Children.SpendingTablesContainer = spendingTablesContainer

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
		savingsTab.UpdateAvailableSavingAmount(availableAmount)
		availableAmount.Refresh()
	}
	amountEntry.Validator = amountEntryValidator

	amountEntryContainer := container.NewBorder(nil, nil, amountText, nil, amountEntry)

	// Save button
	saveBtn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		savingsTab.ValidateAndUpdateSavingAmount(amountEntry.Text)
	})
	saveBtn.Alignment = widget.ButtonAlignTrailing
	saveBtn.Importance = widget.HighImportance

	// Calculate button
	calcBtn := widget.NewButtonWithIcon("Calculate", theme.ViewRefreshIcon(), func() {
		savingsTab.UpdateAvailableSavingAmount(availableAmount)
	})

	savingsContainer := container.NewVBox(
		amountEntryContainer,
		container.NewHBox(
			availableAmount,
			layout.NewSpacer(),
			container.NewHBox(
				calcBtn,
				saveBtn,
			),
		),
	)

	return savingsContainer
}
