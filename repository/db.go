package repository

import (
	"database/sql"
	"errors"
)

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

/**
 * Methods for Savings
 */
func (repo *SQLiteRepository) MigrateSavings() error {
	query := `
		create table if not exists savings(
			id integer primary key autoincrement,
			amount real not null,
			available_amount integer not null,
			year integer
		);
	`

	insertRowWhenNotExists := `
		insert or ignore into savings(amount, available_amount, year)
		select 0, 0, strftime('%Y', 'now')
		where (select count(*) from savings) = 0;
	`

	_, err := repo.Conn.Exec(query)
	if err != nil {
		return err
	}

	_, err = repo.Conn.Exec(insertRowWhenNotExists)
	return err
}

func (s *SQLiteRepository) GetSaving() (*Saving, error) {
	query := `select * from savings;`

	rows, err := s.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	var savings []Saving
	for rows.Next() {
		var saving Saving
		err = rows.Scan(&saving.ID, &saving.Amount, &saving.AvailableAmount, &saving.Year)
		if err != nil {
			return nil, err
		}
		savings = append(savings, saving)
	}

	if len(savings) == 0 {
		return nil, errors.New("no savings found")
	}

	return &savings[0], nil
}

func (s *SQLiteRepository) UpdateSavingAmount(amount int64) error {
	_, err := s.Conn.Exec("update savings set amount = ? where id = 1", amount) // Think better solution for id = 1
	if err != nil {
		return err
	}
	return nil
}

/**
 * Methods for Spendings
 */
func (s *SQLiteRepository) MigrateSpendings() error {
	return nil
}
