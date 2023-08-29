package handler

import (
	"context"
	"net/http"
	"segment-service/internal/core/domain"
	"segment-service/internal/lib/validator"
	"time"

	"github.com/gin-gonic/gin"
)

type segmentServicer interface {
	AddSegment(ctx context.Context, dto *domain.SegmentAddDTO) (*domain.Segment, error)
	DeleteSegment(ctx context.Context, dto *domain.SegmentName) error
	GetAllSegments(ctx context.Context) (*[]domain.Segment, error)
	GetSegmentByName(ctx context.Context, dto *domain.SegmentName) (*domain.Segment, error)
	UpdateSegment(ctx context.Context, dto *domain.SegmentUpdateDTO) (*domain.Segment, error)
	CheckSegmentsExists(ctx context.Context, dto *domain.SegmentNames) error
}

type segmentHandler struct {
	service   segmentServicer
	timeLimit time.Duration
}

func NewSegmentHandler(s segmentServicer) *segmentHandler {
	return &segmentHandler{
		service:   s,
		timeLimit: time.Duration(time.Second * 10),
	}
}

// AddSegment
// @Summary add a segment
// @Tags segment
// @Description add a segment to the database
// @ID get-segment
// @Accept json
// @Produce json
// @Param input body domain.SegmentAddDTO true "segment name"
// @Success 200 {object} domain.Segment
// @Failure 400 {object} badResponse
// @Router /segment [post]
func (h segmentHandler) AddSegment(c *gin.Context) {
	request := &domain.SegmentAddDTO{}
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

	result, err := h.service.AddSegment(ctx, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteSegment
// @Summary delete a segment
// @Tags segment
// @Description delete a segment from the database
// @ID del-segmet
// @Accept json
// @Produce json
// @Param input body domain.SegmentName true "segment name"
// @Success 200 {object} domain.Segment
// @Failure 400 {object} badResponse
// @Router /segment [delete]
func (h segmentHandler) DeleteSegment(c *gin.Context) {
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

	err = h.service.DeleteSegment(ctx, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]bool{"success": true})
}

// GetAllSegments
// @Summary get all segments
// @Tags segment
// @Description return all segments
// @ID get-segmets
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Segment
// @Failure 400 {object} badResponse
// @Router /segment/all [get]
func (h segmentHandler) GetAllSegments(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	result, err := h.service.GetAllSegments(ctx)

	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// UpdateSegment
// @Summary rename a segment
// @Tags segment
// @Description rename a segment
// @ID update-segmets
// @Accept json
// @Produce json
// @Param input body domain.SegmentUpdateDTO true "id and new name of a segment"
// @Success 200 {object} domain.Segment
// @Failure 400 {object} badResponse
// @Router /segment [put]
func (h segmentHandler) UpdateSegment(c *gin.Context) {
	request := &domain.SegmentUpdateDTO{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, badResponse{Message: "bad request"})
		return
	}
	err = validator.ValidateSegmentName(request.NewName)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}
	err = validator.ValidateId(request.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.timeLimit)
	defer cancel()

	result, err := h.service.UpdateSegment(ctx, request)

	if err != nil {
		c.JSON(http.StatusNotFound, badResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
