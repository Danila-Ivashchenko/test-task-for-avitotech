package storage

import (
	"context"
	"segment-service/internal/core/domain"
)

type SegmentStorager interface {
	AddSegment(context.Context, *domain.SegmentAddDTO) (*domain.Segment, error)
	UpdateSegment(context.Context, *domain.SegmentUpdateDTO) (*domain.Segment, error)
	DeleteSegment(context.Context, *domain.SegmentName) error
	GetAllSegments(context.Context) (*[]domain.Segment, error)
	// GetSegmentById(context.Context, *domain.SegmentId) (*domain.Segment, error)
	GetSegmentByName(context.Context, *domain.SegmentName) (*domain.Segment, error)
	GetSegmentsIds(ctx context.Context, dto *domain.SegmentNames) (*domain.SegmentIds, error)

	CheckSegmentsExists(context.Context, *domain.SegmentNames) error
}
