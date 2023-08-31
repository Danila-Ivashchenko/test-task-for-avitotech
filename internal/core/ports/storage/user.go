package storage

import (
	"context"
	"segment-service/internal/core/domain"
)

type UserStorager interface {
	AddUsers(context.Context, *domain.UsersIds) (*domain.UserAffected, error)
	CheckUsersExist(context.Context, *domain.UsersIds) error

	GetUser(context.Context, *domain.UserId) (*domain.User, error)
	GetUsersIds(context.Context, *domain.LinitOffset) (*domain.UsersIds, error)
	GetPercentOfUsersIds(context.Context, *domain.UsersGetPercentDTO) (*domain.UsersIds, error)

	DeleteUsers(context.Context, *domain.UsersIds) (*domain.UserAffected, error)
}
