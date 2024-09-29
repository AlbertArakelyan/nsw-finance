package repository

import "database/sql"

type SQLiteRepository struct {
	Conn *sql.DB
}
type Repository interface {
	MigrateSavings() error
	GetSaving() (*Saving, error)

	MigrateSpendings() error
}

type Saving struct {
	ID              int64   `json:"id"`
	Amount          int64 `json:"amount"`
	AvailableAmount int64 `json:"available_amount"`
	Year            int64   `json:"year"`
}

type Spending struct {
	ID       int64   `json:"id"`
	Amount   float64 `json:"amount"`
	Label    string  `json:"label"`
	Icon     string  `json:"icon"`
	SavingId int64   `json:"saving_id"`
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}
