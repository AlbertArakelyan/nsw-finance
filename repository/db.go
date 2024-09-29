package repository

func (repo *SQLiteRepository) MigrateSavings() error {
	query := `
		create table if not exists savings(
			id integer primary key autoincrement,
			amount real not null,
			available_amount integer not null,
			year integer
		);
	`

	_, err := repo.Conn.Exec(query)
	return err
}

func (s *SQLiteRepository) MigrateSpendings() error {
	return nil
}
