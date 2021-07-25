package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	// AbortWithStatusJSON() блокирует цепочку последующих оброботчиков ендпоинта или маршрута а так же записывает в ответ statusCode и тело сообщения в формате JSON
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
