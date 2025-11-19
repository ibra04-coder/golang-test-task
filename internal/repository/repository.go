package repository

import (
	"database/sql"
)

type NumberRepository interface {
	Save(n int) error
	GetAllSorted() ([]int, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Save(n int) error {
	_, err := r.db.Exec("INSERT INTO numbers (val) VALUES ($1)", n)
	return err
}

func (r *PostgresRepository) GetAllSorted() ([]int, error) {
	rows, err := r.db.Query("SELECT val FROM numbers ORDER BY val ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var numbers []int
	for rows.Next() {
		var n int
		if err := rows.Scan(&n); err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}