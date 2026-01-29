package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Обработчик для скачиваний
func (h *Handler) DownloadFileHandler(c *gin.Context) {
	// Получаем ID
	fileID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Ошибка парсинга fileID: %v", err)
		newErrorResponse(c, http.StatusBadRequest, "invalid file id")
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		log.Printf("Ошибка получения userId: %v", err)
		newErrorResponse(c, http.StatusUnauthorized, "invalid user id")
		return
	}

	log.Printf("Скачивание: userId=%d, fileID=%d", userId, fileID)

	// Получаем информацию о файле из базы данных
	file, err := h.services.File.GetByID(userId, fileID)
	if err != nil {
		log.Printf("Файл не найден в БД: userId=%d, fileID=%d, ошибка: %v", userId, fileID, err)
		newErrorResponse(c, http.StatusNotFound, "file not found")
		return
	}

	log.Printf("Файл из БД: ID=%d, Filename=%s, Path=%s, Size=%d",
		file.ID, file.Filename, file.Path, file.Size)

	// Проверяем существование файла
	if _, err := os.Stat(file.Path); os.IsNotExist(err) {
		log.Printf("Файл не найден на диске: %s", file.Path)
		newErrorResponse(c, http.StatusNotFound, "Файл не найден на диске")
		return
	}

	log.Printf("Отправляем файл: %s", file.Path)

	// Упрощенные заголовки - попробуй так
	c.Header("Content-Disposition", `attachment; filename="`+file.Filename+`"`)
	c.Header("Content-Type", "application/octet-stream")

	c.File(file.Path)
}
