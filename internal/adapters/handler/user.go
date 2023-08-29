package handler

import (
	"context"
	"net/http"
	"segment-service/internal/core/domain"
	"segment-service/internal/lib/validator"
	"time"

	"github.com/gin-gonic/gin"
)

type userServicer interface {
	AddUsers(ctx context.Context, dto *domain.UsersIds) (*domain.UserAffected, error)
	DeleteUsers(ctx context.Context, dto *domain.UsersIds) (*domain.UserAffected, error)
	GetPercentOfUsersIds(ctx context.Context, dto *domain.UsersGetPercentDTO) (*domain.UsersIds, error)
	GetUser(ctx context.Context, dto *domain.UserId) (*domain.User, error)
	GetUsersIds(ctx context.Context, dto *domain.LinitOffset) (*domain.UsersIds, error)
}

type userHandler struct {
	service   userServicer
	timeLimit time.Duration
}

func NewUserHandler(s userServicer) *userHandler {
	return &userHandler{
		service:   s,
		timeLimit: time.Duration(time.Second * 10),
	}
}

// AddUsers
// @Summary add users
// @Tags user
// @Description if some users are registered? they will be ignored
// @ID add-users
// @Accept json
// @Produce json
// @Param input body domain.UsersIds true "users ids"
// @Success 200 {object} domain.UserAffected
// @Failure 400 {object} badResponse
// @Router /user [post]
func (h userHandler) AddUsers(c *gin.Context) {
	request := &domain.UsersIds{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	err = validator.ValidateIds(request.Ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	result, err := h.service.AddUsers(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteUsers
// @Summary delete users
// @Tags user
// @Description if some users aren't registered they will be ignored
// @ID delete-users
// @Accept json
// @Produce json
// @Param input body domain.UsersIds true "users ids"
// @Success 200 {object} domain.UserAffected
// @Failure 400 {object} badResponse
// @Router /user [delete]
func (h userHandler) DeleteUsers(c *gin.Context) {
	request := &domain.UsersIds{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	err = validator.ValidateIds(request.Ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	result, err := h.service.DeleteUsers(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, badResponse{Message: err.Error()})
		return
	}
	responce := map[string]interface{}{
		"deleted": result.Affected,
		"ignored": result.Ignored,
	}
	c.JSON(http.StatusOK, responce)
}

// GetPercentOfUsers
// @Summary get percent of users
// @Tags user
// @Description return percent of random users
// @ID get-percent-users
// @Accept json
// @Produce json
// @Param input body domain.UsersGetPercentDTO true "percent of users"
// @Success 200 {object} domain.UsersIds
// @Failure 400 {object} badResponse
// @Router /user/get/percent [post]
func (h userHandler) GetPercentOfUsersIds(c *gin.Context) {
	request := &domain.UsersGetPercentDTO{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	err = validator.ValidatePercent(request.Percent)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	result, err := h.service.GetPercentOfUsersIds(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetUsersIds
// @Summary get users with limit offset
// @Tags user
// @Description return users with limit, offset, random params
// @ID get-users
// @Accept json
// @Produce json
// @Param input body domain.LinitOffset true "limit offset of users"
// @Success 200 {object} domain.UsersIds
// @Failure 400 {object} badResponse
// @Router /user/get [post]
func (h userHandler) GetUsersIds(c *gin.Context) {
	request := &domain.LinitOffset{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	err = validator.ValidateLimitOffset(request.Limit, request.Offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	result, err := h.service.GetUsersIds(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
