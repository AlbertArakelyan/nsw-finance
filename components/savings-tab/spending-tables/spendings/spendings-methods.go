package spendings

import (
	"nsw-finance/repository"
	"strconv"
)

func (spendings *Spendings) AddNewSpending(savingTableId int64) error {
	err := spendings.DB.AddSpending(savingTableId)
	if err != nil {
		return err
	}

	return nil
}

func (spendings *Spendings) GetSpendings(savingTableId int64) ([]repository.Spending, error) {
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
