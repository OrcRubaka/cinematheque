package repository

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"time"
)

type Movie struct {
	ID          int
	Title       string
	Description string
	ReleaseDate time.Time
	rating      float64
}

type мovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *мovieRepository {
	return &мovieRepository{db: db}
}

// Метод Create создает запись о фильме в базе данных.
func (r *мovieRepository) Create(title, description string, releaseDate time.Time, rating float64) error {
	query := squirrel.Insert("movies").
		Columns("title", "description", "release_date", "rating").
		Values(title, description, releaseDate, rating).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(sqlQuery, args...)
	return err
}

// Метод Get получает информацию о фильме по ID.
func (r *мovieRepository) Get(id int) (*Movie, error) {
	query := squirrel.Select("id", "title", "description", "release_date", "rating").
		From("movies").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(sqlQuery, args...)
	var movie Movie
	err = row.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.rating)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

// Метод Update обновляет информацию о фильме по ID.
func (r *мovieRepository) Update(id int, title, description string, releaseDate time.Time, rating float64) error {
	query := squirrel.Update("movies").
		Set("title", title).
		Set("description", description).
		Set("release_date", releaseDate).
		Set("rating", rating).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(sqlQuery, args...)
	return err
}

// Метод Delete удаляет фильм по ID.
func (r *мovieRepository) Delete(id int) error {
	query := squirrel.Delete("movies").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sqlQuery, args...)
	return err
}
