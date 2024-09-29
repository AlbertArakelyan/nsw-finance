package repository

import "database/sql"

type SQLiteRepository struct {
	Conn *sql.DB
}
type Repository interface {
	MigrateSavings() error
	MigrateSpendings() error
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
