package spendings

import "nsw-finance/repository"

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
