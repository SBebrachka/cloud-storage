package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sbebrachka/pet3/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Раздача фронтенда
	router.Static("/static", "./static")

	// Загружаем HTML файлы
	router.LoadHTMLFiles(
		"static/index.html",
		"static/login.html",
		"static/register.html",
		"static/home.html",
	)

	// Маршруты HTML
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	router.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})
	router.GET("/home", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	// Аутентификация
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// API с middleware аутентификации
	api := router.Group("/api", h.userIdentity)
	{
		// Работа со списками
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			// Элементы списков
			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItem)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}

		// Работа с файлами
		files := api.Group("/files")
		{
			files.GET("/", h.getAllFiles)
			files.POST("/upload", h.uploadFile)
			files.DELETE("/:id", h.deleteFile)
			files.GET("/:id/download", h.DownloadFileHandler)
			files.GET("/:id", h.GetFileByID)
		}
	}

	return router
}
