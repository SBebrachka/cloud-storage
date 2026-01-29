package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sbebrachka/pet3/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Константы для контекста
const (
	userCtx = "userId"
)

// Структуры ответов
type NewErrorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// Вспомогательные функции для ответов
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf("HTTP %d: %s", statusCode, message)
	c.AbortWithStatusJSON(statusCode, NewErrorResponse{Message: message})
}

// Вспомогательная функция для получения ID пользователя из контекста
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.ErrUserNotFound
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.ErrInvalidUserIdType
	}

	return idInt, nil
}
