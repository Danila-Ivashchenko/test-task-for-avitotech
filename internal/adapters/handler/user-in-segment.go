package handler

import (
	"context"
	"net/http"
	"segment-service/internal/core/domain"
	"segment-service/internal/lib/validator"
	"time"

	"github.com/gin-gonic/gin"
)

type userInSegmentServicer interface {
	AddUserToSegments(ctx context.Context, dto *domain.UserToSegmentsAddDTO) error
	AddUsersToSegments(ctx context.Context, dto *domain.UsersToSegmentsAddDTO) error
	AddUsersWithLimitOffsetToSegments(ctx context.Context, dto *domain.UsersWithLimitOffsetToSegments) error
	AddPercentOfUsersToSegments(ctx context.Context, dto *domain.PercentOfUsersToSegmentsDTO) error

	DeleteUserFromSegments(ctx context.Context, dto *domain.UserFromSegmentsDeleteDTO) error
	GetUserInSegments(ctx context.Context, dto *domain.UserId) (*domain.UserInSegments, error)
	GetUsersInSegment(ctx context.Context, dto *domain.SegmentName) (*domain.UsersInSegment, error)
}

type userInSegmentHandler struct {
	service   userInSegmentServicer
	timeLimit time.Duration
}

func NewUserInSegmentHandler(s userInSegmentServicer) *userInSegmentHandler {
	return &userInSegmentHandler{
		service:   s,
		timeLimit: time.Duration(time.Second * 10),
	}
}

// AddUsersToSegments
// @Summary bind users and segments
// @Tags user-in-segment
// @Description bind users by ids and segments by names
// @ID add-users-to-segments
// @Accept json
// @Produce json
// @Param input body domain.UsersToSegmentsAddDTO true "users ids and segments names"
// @Success 200 {object} successResponse
// @Failure 400 {object} badResponse
// @Router /user_in_segment/add/many [post]
func (h userInSegmentHandler) AddUsersToSegments(c *gin.Context) {
	request := &domain.UsersToSegmentsAddDTO{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: "bad request"})
		return
	}
	err = validator.ValidateSegmentNames(request.SegmentNames)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}
	err = validator.ValidateIds(request.UserIds)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	err = h.service.AddUsersToSegments(ctx, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, successResponse{Success: true})
}

// AddUsersWithLimitOffsetToSegments
// @Summary bind users and segments
// @Tags user-in-segment
// @Description bind users by limit/offset and segments by names
// @ID add-users-limit-offset-to-segments
// @Accept json
// @Produce json
// @Param input body domain.UsersWithLimitOffsetToSegments true "limit, offset, random for getting users and segments names"
// @Success 200 {object} successResponse
// @Failure 400 {object} badResponse
// @Router /user_in_segment/add/params [post]
func (h userInSegmentHandler) AddUsersWithLimitOffsetToSegments(c *gin.Context) {
	request := &domain.UsersWithLimitOffsetToSegments{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: "bad request"})
		return
	}
	err = validator.ValidateSegmentNames(request.SegmentNames)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}
	err = validator.ValidateLimitOffset(request.Limit, request.Offset)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	err = h.service.AddUsersWithLimitOffsetToSegments(ctx, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, successResponse{Success: true})
}

// AddPercentOfUsersToSegments
// @Summary bind users and segments
// @Tags user-in-segment
// @Description bind users by persent and segments by names
// @ID add-users-persent-to-segments
// @Accept json
// @Produce json
// @Param input body domain.PercentOfUsersToSegmentsDTO true "persent of users and segments names"
// @Success 200 {object} successResponse
// @Failure 400 {object} badResponse
// @Router /user_in_segment/add/persent [post]
func (h userInSegmentHandler) AddPercentOfUsersToSegments(c *gin.Context) {
	request := &domain.PercentOfUsersToSegmentsDTO{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: "bad request"})
		return
	}
	err = validator.ValidateSegmentNames(request.SegmentNames)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}
	err = validator.ValidatePercent(request.Percent)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	err = h.service.AddPercentOfUsersToSegments(ctx, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, successResponse{Success: true})
}

// AddUserToSegments
// @Summary bind user and segments
// @Tags user-in-segment
// @Description bind user by id and segments by names
// @ID add-user-to-segments
// @Accept json
// @Produce json
// @Param input body domain.UserToSegmentsAddDTO true "user id and segments names"
// @Success 200 {object} successResponse
// @Failure 400 {object} badResponse
// @Router /user_in_segment/add/one [post]
func (h userInSegmentHandler) AddUserToSegments(c *gin.Context) {
	request := &domain.UserToSegmentsAddDTO{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: "bad request"})
		return
	}
	err = validator.ValidateSegmentNames(request.SegmentNames)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}
	err = validator.ValidateId(request.UserId)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	err = h.service.AddUserToSegments(ctx, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, successResponse{Success: true})
}

// DeleteUserFromSegments
// @Summary delete users from segments
// @Tags user-in-segment
// @Description delete users from segments by names
// @ID delete-users-from-segments
// @Accept json
// @Produce json
// @Param input body domain.UserFromSegmentsDeleteDTO true "users ids and segments names"
// @Success 200 {object} successResponse
// @Failure 400 {object} badResponse
// @Router /user_in_segment [delete]
func (h userInSegmentHandler) DeleteUserFromSegments(c *gin.Context) {
	request := &domain.UserFromSegmentsDeleteDTO{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: "bad request"})
		return
	}
	err = validator.ValidateSegmentNames(request.SegmentNames)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}
	err = validator.ValidateId(request.UserId)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	err = h.service.DeleteUserFromSegments(ctx, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]bool{"success": true})
}

// GetUserInSegments
// @Summary get user's segments
// @Tags user-in-segment
// @Description get user's segments
// @ID get-segment-of-user
// @Accept json
// @Produce json
// @Param input body domain.UserId true "users id"
// @Success 200 {object} successResponse
// @Failure 400 {object} badResponse
// @Router /user_in_segment/get/user [post]
func (h userInSegmentHandler) GetUserInSegments(c *gin.Context) {
	request := &domain.UserId{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: "bad request"})
		return
	}
	err = validator.ValidateId(request.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	responce, err := h.service.GetUserInSegments(ctx, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, responce)
}

// GetUserInSegments
// @Summary get users in segments
// @Tags user-in-segment
// @Description get users in segments
// @ID get-users-in-segment
// @Accept json
// @Produce json
// @Param input body domain.SegmentName true "segment name"
// @Success 200 {object} successResponse
// @Failure 400 {object} badResponse
// @Router /user_in_segment/get/segment [post]
func (h userInSegmentHandler) GetUsersInSegment(c *gin.Context) {
	request := &domain.SegmentName{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: "bad request"})
		return
	}
	err = validator.ValidateSegmentName(request.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	responce, err := h.service.GetUsersInSegment(ctx, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, responce)
}
