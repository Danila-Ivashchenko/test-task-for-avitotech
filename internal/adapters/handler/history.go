package handler

import (
	"context"
	"net/http"
	"segment-service/internal/core/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type historyServicer interface {
	GetHistoryOfUser(ctx context.Context, dto *domain.HistoryOfUserGetDTO) (*domain.HistoryResponce, error)
}

type historyHandler struct {
	service   historyServicer
	timeLimit time.Duration
}

func NewHistoryHandler(s historyServicer) *historyHandler {
	return &historyHandler{
		service:   s,
		timeLimit: time.Duration(time.Second * 60),
	}
}

// GetHistoryOfUser
// @Summary return history of a user
// @Tags history
// @Description returl url of a file with user's fistory in csv format
// @ID get-history
// @Accept json
// @Produce json
// @Param input body domain.HistoryOfUserGetDTO true "user_id, month, year"
// @Success 200 {object} domain.HistoryResponce
// @Failure 400 {object} badResponse
// @Router /history [post]
func (h historyHandler) GetHistoryOfUser(c *gin.Context) {
	request := &domain.HistoryOfUserGetDTO{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	result, err := h.service.GetHistoryOfUser(ctx, request)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
