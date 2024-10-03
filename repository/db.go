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

func (repo *SQLiteRepository) GetSaving() (*Saving, error) {
	query := `select * from savings;`

	rows, err := repo.Conn.Query(query)
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
func (repo *SQLiteRepository) MigrateSpendingTables() error {
	query := `
		create table if not exists spending_tables(
			id integer primary key autoincrement,
			label text not null,
			saving_id integer not null,

			foreign key(saving_id) references savings(id)
		);
	`

	_, err := repo.Conn.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (repo *SQLiteRepository) MigrateSpendings() error {
	return nil
}

func (repo *SQLiteRepository) AddSpendingTable(label string, savingId int64) (int64, error) {
	res, err := repo.Conn.Exec("insert into spending_tables(label, saving_id) values(?, ?)", label, savingId)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (repo *SQLiteRepository) GetSpendingTables(savingId int64) ([]SpendingTable, error) {
	query := `select * from spending_tables where saving_id = ?;`

	rows, err := repo.Conn.Query(query, savingId)
	if err != nil {
		return nil, err
	}

	var spendingTables []SpendingTable
	for rows.Next() {
		var spendingTable SpendingTable
		err = rows.Scan(&spendingTable.ID, &spendingTable.Label, &spendingTable.SavingId)
		if err != nil {
			return nil, err
		}
		spendingTables = append(spendingTables, spendingTable)
	}

	return spendingTables, nil
}

func (repo *SQLiteRepository) GetSpendingTableByID(id int64) (*SpendingTable, error) {
	query := `select * from spending_tables where id = ?;`

	rows, err := repo.Conn.Query(query, id)
	if err != nil {
		return nil, err
	}

	var spendingTable SpendingTable
	for rows.Next() {
		err = rows.Scan(&spendingTable.ID, &spendingTable.Label, &spendingTable.SavingId)
		if err != nil {
			return nil, err
		}
	}

	return &spendingTable, nil
}
