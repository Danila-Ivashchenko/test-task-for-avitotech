package history

import (
	"context"
	"fmt"
	"segment-service/internal/adapters/storage/lib"
	"segment-service/internal/core/domain"
)

type historyStorage struct {
	manager lib.PsqlManager
}

func NewHistoryStorage(m lib.PsqlManager) *historyStorage {
	return &historyStorage{
		manager: m,
	}
}

func (s historyStorage) getSegmentIds(dto *domain.SegmentNames) ([]int64, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()

	var (
		id         int64
		segmentIds = make([]int64, 0, len(dto.Names))
	)

	stmt := fmt.Sprintf("SELECT id FROM segments WHERE name IN%s", lib.ArrayStrToStr(dto.Names))
	rows, err := db.Queryx(stmt)
	if err != nil {
		return segmentIds, domain.ErrorNoSuchSegments
	}

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return segmentIds, err
		}
		segmentIds = append(segmentIds, id)
	}
	return segmentIds, nil
}

func (s historyStorage) AddHistory(dto *domain.HistoryAddDTO) {
	db, err := s.manager.GetDb()
	if err != nil {
		return
	}
	defer db.Close()

	segmentIds, err := s.getSegmentIds(&domain.SegmentNames{Names: dto.SegmentNames})
	if err != nil {
		return
	}

	stmt := fmt.Sprintf("INSERT INTO history (user_id, segment_id, action) VALUES %s", lib.MakeHistoryRecords(dto.UserIds, segmentIds, dto.Action))
	fmt.Println(stmt)
	db.Exec(stmt)
}

func (s historyStorage) GetHistoryOfUser(ctx context.Context, dto *domain.HistoryOfUserGetDTO) (*[]domain.HistoryOfUser, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()

	stmt := "SELECT segments.name, history.action, history.action_time FROM history, segments"
	stmt += fmt.Sprintf(" WHERE history.segment_id = segments.id AND history.user_id = %d AND EXTRACT(MONTH FROM history.action_time) = %d AND EXTRACT(YEAR FROM history.action_time) = %d", dto.UserId, dto.Month, dto.Year)

	rows, err := db.QueryxContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	record := domain.HistoryOfUser{}
	records := make([]domain.HistoryOfUser, 0)
	for rows.Next() {
		err := rows.StructScan(&record)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return &records, nil
}
