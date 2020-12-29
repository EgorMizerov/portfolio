package handler

import (
	"github.com/EgorMizerov/portfolio/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getAllWorksResponse struct {
	Data []models.Work `json:"data"`
}

func (h *Handler) getWorks(c *gin.Context) {
	lists, err := h.service.Work.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllWorksResponse{
		Data: lists,
	})
}

func (h *Handler) createWork(c *gin.Context) {
	var input models.Work

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Work.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Work Created",
		"id":      id,
	})
}

func (h *Handler) deleteWork(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.service.Work.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func (h *Handler) updateWork(c *gin.Context) {

}
