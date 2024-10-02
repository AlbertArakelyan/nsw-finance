package savingstab

import (
	"fmt"
	"log"
	"nsw-finance/repository"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SavingsTab struct {
	DB       repository.Repository
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var currency = "AMD"

func (savingsTab *SavingsTab) GetSavingsTab() *fyne.Container {
	savingsTextContainer := savingsTab.getSavingsTextContainer()
	savingsContainer := container.NewVBox(savingsTextContainer)

	return savingsContainer
}

func (savingsTab *SavingsTab) getSavingsTextContainer() *fyne.Container {
	amount, availableAmount := savingsTab.getSavingsText()

	amountText := canvas.NewText("Savings: ", nil)
	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder(fmt.Sprintf("Amount (%s)", currency))
	amountEntry.Text = amount
	amountEntry.Validator = func(s string) error {
		_, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		return nil
	}

	amountEntryContainer := container.NewBorder(nil, nil, amountText, nil, amountEntry)

	saveBtn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		err := savingsTab.UpdateSavingAmount(amountEntry.Text)
		if err != nil {
			savingsTab.ErrorLog.Println(err)
			log.Panic(err)
			return
		}

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
