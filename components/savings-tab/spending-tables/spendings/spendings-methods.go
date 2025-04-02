package spendings

import (
	savingsrepository "nsw-finance/repository/savings-repository"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (spendings *Spendings) AddNewSpending(savingTableId int64, spendingsList *fyne.Container) error {
	s, err := spendings.DB.AddSpending(savingTableId)
	if err != nil {
		return err
	}

	var spendingsSlice []savingsrepository.Spending

	spendingsSlice = append(spendingsSlice, *s)

	spendings.appendSpendingToList(spendingsList, spendingsSlice)

	return nil
}

func (spendings *Spendings) GetSpendings(savingTableId int64) ([]savingsrepository.Spending, error) {
	spendingsSlice, err := spendings.DB.GetSpendings(savingTableId)
	if err != nil {
		return nil, err
	}

	return spendingsSlice, nil
}

func (spendings *Spendings) spendingAmountValidator(s string) error {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	return nil
}

func (spendings *Spendings) ValidateAndUpdateSpendingAmount(spendingId int64, amountText string) error {
	err := spendings.spendingAmountValidator(amountText)
	if err != nil {
		return err
	}

	amount, err := strconv.ParseFloat(amountText, 64)
	if err != nil {
		return err
	}

	err = spendings.DB.UpdateSpendingAmount(spendingId, amount)
	if err != nil {
		return err
	}

	return nil
}

func (spendings *Spendings) UpdateSpendingLabel(id int64, label string) error {
	err := spendings.DB.UpdateSpendingLabel(id, label)
	if err != nil {
		return err
	}

	return nil
}

func (spendings *Spendings) appendSpendingToList(spendingsList *fyne.Container, spendingsSlice []savingsrepository.Spending) {
	for _, spending := range spendingsSlice {
		spendingLabelEntry := widget.NewEntry()
		spendingLabelEntry.SetText(spending.Label)
		spendingLabelEntry.OnChanged = func(s string) {
			if s == "" {
				s = "New Spending"
				spendingLabelEntry.SetText(s)
			}
			spendingLabelEntry.OnChanged = func(s string) {
				if s == "" {
					s = "New Spending"
					spendingLabelEntry.SetText(s)
				}
				go func() {
					time.Sleep(500 * time.Millisecond)
					err := spendings.UpdateSpendingLabel(spending.ID, s)
					if err != nil {
						spendings.ErrorLog.Println(err)
						// log.Panic(err)
					}
				}()
			}
		}

		amountText := strconv.FormatFloat(float64(spending.Amount), 'f', 2, 64)
		spendingAmountEntry := widget.NewEntry()
		spendingAmountEntry.SetText(amountText)
		spendingAmountEntry.Validator = spendings.spendingAmountValidator
		spendingAmountEntry.OnChanged = func(s string) {
			err := spendings.ValidateAndUpdateSpendingAmount(spending.ID, s)
			if err != nil {
				spendings.ErrorLog.Println(err)
				// log.Panic(err)
			}

			// TODO update available amount on every change
		}

		c := container.NewGridWithColumns(
			2,
			spendingLabelEntry,
			spendingAmountEntry,
		)

		spendingsList.Add(c)
	}
}
