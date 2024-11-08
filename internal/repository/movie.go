package repository

import "database/sql"

type movie struct {
	db *sql.DB
}

func NewMovie(db *sql.DB) *movie {
	return &movie{db: db}
}

// Создание нового фильма
func (m *movie) Create(title string, description string, releaseDate string, rating float64) error {
	query := `INSERT INTO movies (title, description, release_date, rating) VALUES ($1, $2, $3, $4)`
	_, err := m.db.Exec(query, title, description, releaseDate, rating)
	return err
}

// Редактирование информации о фильме
func (m *movie) Update(id int, title string, description string, releaseDate string, rating float64) error {
	query := `UPDATE movies SET title=$1, description=$2, release_date=$3, rating=$4 WHERE id=$5`
	_, err := m.db.Exec(query, title, description, releaseDate, rating, id)
	return err
}
