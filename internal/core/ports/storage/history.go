package storage

import (
	"segment-service/internal/core/domain"

	"golang.org/x/net/context"
)

type HistoryStorage interface {
	AddHistory(dto *domain.HistoryAddDTO)
	GetHistoryOfUser(context.Context, *domain.HistoryOfUserGetDTO) (*[]domain.HistoryOfUser, error)
}
