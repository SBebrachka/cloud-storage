package service

import (
	"fmt"
	"github.com/sbebrachka/pet3/pkg/model" // ИЛИ pkg/model
	"github.com/sbebrachka/pet3/pkg/repository"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type FileService struct {
	repo repository.File
}

func NewFileService(repo repository.File) *FileService {
	return &FileService{repo: repo}
}

func (s *FileService) GetAll(userId int) ([]model.File, error) {
	return s.repo.GetAll(userId)
}

func (s *FileService) Upload(userId int, filename string, size int64, file io.Reader) (model.File, error) {
	// Преобразуем userId в строку для пути
	userDir := filepath.Join("uploads", strconv.Itoa(userId))

	// Создаем директорию для файлов пользователя, если ее нет
	if err := os.MkdirAll(userDir, 0755); err != nil {
		return model.File{}, fmt.Errorf("failed to create directory: %w", err)
	}

	// Создаем путь к файлу
	filePath := filepath.Join(userDir, filename)

	// Создаем файл на диске
	dst, err := os.Create(filePath)
	if err != nil {
		return model.File{}, fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	// Копируем содержимое
	if _, err := io.Copy(dst, file); err != nil {
		// Если копирование не удалось, удаляем созданный файл
		os.Remove(filePath)
		return model.File{}, fmt.Errorf("failed to copy file content: %w", err)
	}

	// Сохраняем информацию о файле в БД
	return s.repo.Create(userId, model.File{
		Filename: filename,
		Size:     size,
		Path:     filePath,
	})
}

func (s *FileService) Delete(userId, fileId int) error {
	// Получаем информацию о файле
	file, err := s.repo.GetByID(userId, fileId)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	// Удаляем файл с диска
	if err := os.Remove(file.Path); err != nil && !os.IsNotExist(err) {
		// Игнорируем ошибку, если файл уже не существует
		return fmt.Errorf("failed to delete file from disk: %w", err)
	}

	// Удаляем запись из БД
	if err := s.repo.Delete(userId, fileId); err != nil {
		return fmt.Errorf("failed to delete file record from DB: %w", err)
	}

	return nil
}

func (s *FileService) GetByID(userId, id int) (model.File, error) {
	return s.repo.GetByID(userId, id)
}

func (s *FileService) Search(userId int, query string) ([]model.File, error) {
	return s.repo.Search(userId, query)
}
