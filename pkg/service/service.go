package service

import (
	"io"

	"github.com/sbebrachka/pet3"
	"github.com/sbebrachka/pet3/pkg/model"
	"github.com/sbebrachka/pet3/pkg/repository"
)

// Интерфейсы
type Authorization interface {
	CreateUser(user pet3.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type BoxList interface {
	// Методы для работы со списками
}

type BoxItem interface {
	// Методы для работы с элементами
}

type File interface {
	GetAll(userID int) ([]model.File, error)
	Upload(userID int, filename string, size int64, file io.Reader) (model.File, error)
	Delete(userID, fileID int) error
	GetByID(userID, fileID int) (model.File, error)
}

// Основная структура сервиса
type Service struct {
	Authorization
	BoxList
	BoxItem
	File
}

// Конструктор
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		File:          NewFileService(repos.File),
	}
}
