package storage

import (
	"context"
	"segment-service/internal/core/domain"
)

type UserInSegmentStorager interface {
	AddUserToSegments(context.Context, *domain.UserToSegmentsAddDTO) error
	AddUsersToSegments(context.Context, *domain.UsersToSegmentsAddDTO) error
	AddPercentOfUsersToSegments(context.Context, *domain.PercentOfUsersToSegmentsDTO) (*domain.UsersIds, error)
	AddUsersWithLimitOffsetToSegments(context.Context, *domain.UsersWithLimitOffsetToSegments) (*domain.UsersIds, error)
	DeleteUserFromSegments(context.Context, *domain.UserFromSegmentsDeleteDTO) error
	GetUserInSegments(context.Context, *domain.UserId) (*domain.UserInSegments, error)
	GetUsersInSegment(context.Context, *domain.SegmentName) (*domain.UsersInSegment, error)
}
