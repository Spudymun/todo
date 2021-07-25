package handler

import (
	"net/http"

	"github.com/Spudymun/todo"
	"github.com/gin-gonic/gin"
)

// Handler в Фреймворке Gin, это ф-ция которая должна иметь в качестве параметра указатель на обьект gin.Context
func (h *Handler) signUp(c *gin.Context) {
	// Создание структуры input, в которой записываються данные из JSON о пользователей
	var input todo.User

	// c.BindJSON() принимает ссылку на обьект в котором мы хотим распарсить тело JSON и присвоить поля JSON полям струтуры input. Так же будед проведина валидация тега binding
	if err := c.BindJSON(&input); err != nil {
		// Вызов ф-ции для создания ответа с ошибкой. http.StatusBadRequest=400(некоректные данные в запросе)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Вызов метода сервиса CreateUser() в котрый передаеться структура пользователя
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		// Запись в ответStatusInternalServerError=500
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Запись в ответ StatusOK=200 и тело JSON со значением id пользователя
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// Новая структура для авторизации
type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	// Создание структуры input типа signInInput, в которой записываються данные из JSON о пользователей
	var input signInInput

	// c.BindJSON() принимает ссылку на обьект в которой мы хотим распарсить тело JSON и присвоить поля JSON полям струтуры input. Так же будед проведина валидация тега binding
	if err := c.BindJSON(&input); err != nil {
		// Вызов ф-ции для создания ответа с ошибкой. http.StatusBadRequest=400(некоректные данные в запросе)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Вызов метода сервиса GenerateToken() в котрый передаеться поля username и password
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		// Запись в ответStatusInternalServerError=500
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Запись в ответ StatusOK=200 и тело JSON со значением token пользователя
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
