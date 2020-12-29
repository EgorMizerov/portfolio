package handler

import (
	"github.com/EgorMizerov/portfolio/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", h.getIndex)
	api := router.Group("/api")
	{
		api.POST("/", h.createWork)
		api.GET("/", h.getWorks)
		api.PUT("/:id", h.updateWork)
		api.DELETE("/:id", h.deleteWork)
	}

	return router
}
