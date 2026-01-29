package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Функции для работы с файлами
func (h *Handler) getAllFiles(c *gin.Context) {
	// Получаем ID пользователя из контекста
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get user id: "+err.Error())
		return
	}

	// Получаем файлы из сервиса
	files, err := h.services.File.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get files: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, files)
}

func (h *Handler) uploadFile(c *gin.Context) {
	// Получаем ID пользователя из контекста
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get user id: "+err.Error())
		return
	}

	// Получаем файл из запроса
	file, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "file is required: "+err.Error())
		return
	}

	// Проверяем размер файла (максимум 10MB)
	if file.Size > 10<<20 { // 10MB в байтах
		newErrorResponse(c, http.StatusBadRequest, "file size exceeds 10MB limit")
		return
	}

	// Проверяем, что файл не пустой
	if file.Size == 0 {
		newErrorResponse(c, http.StatusBadRequest, "file is empty")
		return
	}

	// Открываем файл
	src, err := file.Open()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to open file: "+err.Error())
		return
	}
	defer src.Close()

	// Сохраняем файл через сервис
	uploadedFile, err := h.services.File.Upload(userId, file.Filename, file.Size, src)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to upload file: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, uploadedFile)
}

func (h *Handler) deleteFile(c *gin.Context) {
	// Получаем ID пользователя из контекста
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get user id: "+err.Error())
		return
	}

	// Получаем ID файла из параметра URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param: "+err.Error())
		return
	}

	// Проверяем что ID положительный
	if id <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid file id")
		return
	}

	// Удаляем файл через сервис
	err = h.services.File.Delete(userId, id)
	if err != nil {
		// Проверяем тип ошибки для более точного HTTP статуса
		if err.Error() == "file not found" || err.Error() == "file not found or permission denied" {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "failed to delete file: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status:  "ok",
		Message: "file deleted successfully",
	})
}

func (h *Handler) GetFileByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get user id: "+err.Error())
		return
	}
	fileID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid file id")
		return
	}
	file, err := h.services.File.GetByID(userId, fileID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "file not found")
		return
	}

	c.JSON(http.StatusOK, file)
}
