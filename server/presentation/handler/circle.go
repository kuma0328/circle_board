package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuma0328/circle_board/domain/entity"
	"github.com/kuma0328/circle_board/usecase"
)

type CircleHandler struct {
	uc usecase.CircleUsecase
}

func NewCircleHandler(uc usecase.CircleUsecase) *CircleHandler {
	return &CircleHandler{
		uc: uc,
	}
}

func (h *CircleHandler) CreateCircle(c *gin.Context) {
	var circle entity.Circle
	if err := c.BindJSON(&circle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if err := h.uc.CreateCircle(&circle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, circle)
}

func (h *CircleHandler) GetAllCircle(c *gin.Context) {
	circles, err := h.uc.GetAllCircle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, circles)
}

func (h *CircleHandler) GetCircleById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	circle, err := h.uc.GetCircleById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, circle)
}