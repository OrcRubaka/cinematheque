package repository

import "database/sql"

type actor struct {
	db *sql.DB
}

func NewActor(db *sql.DB) *actor {
	return &actor{db: db}
}

// Создание нового актера
func (a *actor) Create(name string, gender string, birthDate string) error {
	query := `INSERT INTO actors (name, gender, birth_date) VALUES ($1, $2, $3)`
	_, err := a.db.Exec(query, name, gender, birthDate)
	return err
}

// Редактирование информации об актере
func (a *actor) Update(id int, name string, gender string, birthDate string) error {
	query := `UPDATE actors SET name=$1, gender=$2, birth_date=$3 WHERE id=$4`
	_, err := a.db.Exec(query, name, gender, birthDate, id)
	return err
}
