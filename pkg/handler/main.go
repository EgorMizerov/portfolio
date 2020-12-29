package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getIndex(c *gin.Context) {
	lists, err := h.service.Work.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(len(lists))

	c.HTML(http.StatusOK, "index.html", gin.H{
		"data": lists,
	})
}
