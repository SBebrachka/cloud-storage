package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/sbebrachka/pet3/pkg/model"
)

type FilePostgres struct {
	db *sqlx.DB
}

func NewFilePostgres(db *sqlx.DB) *FilePostgres {
	return &FilePostgres{db: db}
}

func (r *FilePostgres) GetAll(userId int) ([]model.File, error) {
	var files []model.File
	query := `SELECT id, user_id, filename, size, path, created_at FROM files WHERE user_id = $1 ORDER BY created_at DESC`
	err := r.db.Select(&files, query, userId)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (r *FilePostgres) GetByID(userId, fileId int) (model.File, error) {
	var file model.File
	query := `SELECT id, user_id, filename, size, path, created_at FROM files WHERE id = $1 AND user_id = $2`
	err := r.db.Get(&file, query, fileId, userId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return model.File{}, errors.New("file not found")
		}
		return model.File{}, err
	}
	return file, nil
}

func (r *FilePostgres) Create(userId int, file model.File) (model.File, error) {
	var id int
	query := `INSERT INTO files (user_id, filename, size, path) VALUES ($1, $2, $3, $4) RETURNING id, created_at`

	err := r.db.QueryRow(query, userId, file.Filename, file.Size, file.Path).Scan(&id, &file.CreatedAt)
	if err != nil {
		return model.File{}, err
	}

	file.ID = id
	file.UserId = userId
	return file, nil
}

func (r *FilePostgres) Delete(userId, fileId int) error {
	query := `DELETE FROM files WHERE id = $1 AND user_id = $2`
	result, err := r.db.Exec(query, fileId, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("file not found or permission denied")
	}

	return nil
}
func (r *FilePostgres) Search(userId int, search string) ([]model.File, error) {
	var files []model.File

	query := `
		SELECT id, user_id, filename, size, path, created_at
		FROM files
		WHERE user_id = $1 AND filename ILIKE '%' || $2 || '%'
		ORDER BY created_at DESC
	`

	err := r.db.Select(&files, query, userId, search)
	return files, err
}
