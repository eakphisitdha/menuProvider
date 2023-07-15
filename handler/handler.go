package handler

import (
	"encoding/json"
	"menuProvider/model"
	"menuProvider/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Get(c *gin.Context)
}

type Handler struct {
	s service.IService
}

func NewHandler(s service.IService) IHandler {
	return &Handler{s: s}
}

func (h *Handler) Get(c *gin.Context) {
	menu, err := h.s.GetMenu()
	if err != nil {
		response := model.Response{
			Status: false,
			Code:   404,
			Data:   "unfound",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := model.Response{
		Status: true,
		Code:   200,
		Data:   menu,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Data(http.StatusOK, "application/json", jsonResponse)
}
