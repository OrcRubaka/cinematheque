package repository

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"time"
)

type Actor struct {
	ID        int
	Name      string
	Gender    string
	BirthDate time.Time
}

type actorRepository struct {
	db *sql.DB
}

func NewActorRepository(db *sql.DB) *actorRepository {
	return &actorRepository{db: db}
}

// Создание актера
func (r *actorRepository) Create(name string, gender string, birthDate time.Time) error {
	query := squirrel.Insert("actors").
		Columns("name", "gender", "birth_date").
		Values(name, gender, birthDate).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sqlQuery, args...)
	return err
}

// Обновление актера
func (r *actorRepository) Update(id int, name string, birthDate time.Time) error {
	// В запросе обновляется только имя и дата рождения
	query := squirrel.Update("actors").
		Set("name", name).
		Set("birth_date", birthDate).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sqlQuery, args...)
	return err
}

// Получение актера
func (r *actorRepository) Get(id int) (*Actor, error) {
	query := squirrel.Select("id", "name", "gender", "birth_date").
		From("actors").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(sqlQuery, args...)
	var actor Actor
	err = row.Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.BirthDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &actor, nil
}

// Удаление актера
func (r *actorRepository) Delete(id int) error {
	query := squirrel.Delete("actors").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sqlQuery, args...)
	return err
}
