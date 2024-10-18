package savingsrepository

import "database/sql"

type SQLiteRepository struct {
	Conn *sql.DB
}

type Repository interface {
	MigrateSavings() error
	GetSaving() (*Saving, error)
	UpdateSavingAmount(amount int64) error
	UpdateAvailableSavingAmount() (int64, error)

	MigrateSpendingTables() error
	MigrateSpendings() error
	AddSpendingTable(label string, savingId int64) (int64, error)
	GetSpendingTables(savingId int64) ([]SpendingTable, error)
	GetSpendingTableByID(id int64) (*SpendingTable, error)

	AddSpending(savingTableId int64) (*Spending, error)
	GetSpendings(savingTableId int64) ([]Spending, error)
	GetSpendingByID(id int64) (*Spending, error)
	UpdateSpendingAmount(id int64, amount float64) error
	UpdateSpendingLabel(id int64, label string) error
}

type Saving struct {
	ID              int64 `json:"id"`
	Amount          float64 `json:"amount"`
	AvailableAmount int64 `json:"available_amount"`
	Year            int64 `json:"year"`
}

type SpendingTable struct {
	ID       int64  `json:"id"`
	Label    string `json:"label"`
	SavingId int64  `json:"saving_id"`
}

type Spending struct {
	ID              int64   `json:"id"`
	Amount          float64 `json:"amount"`
	Label           string  `json:"label"`
	Icon            *string `json:"icon"`
	SpendingTableId int64   `json:"spending_table_id"`
}
