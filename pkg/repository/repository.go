package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sbebrachka/pet3"
	"github.com/sbebrachka/pet3/pkg/model"
)

// Интерфейсы репозиториев
type Authorization interface {
	CreateUser(user pet3.User) (int, error)
	GetUser(username, password string) (pet3.User, error)
}

type BoxList interface {
	// Методы будут добавлены позже
}

type BoxItem interface {
	// Методы будут добавлены позже
}

type File interface {
	GetAll(userId int) ([]model.File, error)
	GetByID(userId, fileId int) (model.File, error)
	Create(userId int, file model.File) (model.File, error)
	Delete(userId, fileId int) error
	Search(userId int, query string) ([]model.File, error)
}

// Основная структура репозитория
type Repository struct {
	Authorization
	BoxList
	BoxItem
	File
}

// Конструктор
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		BoxList:       nil, // будет реализовано позже
		BoxItem:       nil, // будет реализовано позже
		File:          NewFilePostgres(db),
	}
}
