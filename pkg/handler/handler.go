package handler

import (
	"github.com/Spudymun/todo/pkg/service"
	"github.com/gin-gonic/gin"
)

// Оброботчики будут вызывать методы сервиса, поэтому в структуре указатель на сервисы
type Handler struct {
	services *service.Service
}

// Внедрение зависимотсей от сервиса
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// Функция для инициализации endpoints
func (h *Handler) InitRoutes() *gin.Engine {
	// Для разработки REST API используем WEB/HTTP-фреймворк Gin
	// Для инициализации роутера вызовем Gim.New()
	router := gin.New()

	// Обьявление методов згрупировав их по группам
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	return router
}
