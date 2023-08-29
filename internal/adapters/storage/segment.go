package storage

import (
	"context"
	"fmt"
	"segment-service/internal/core/domain"
)

type segmentStorage struct {
	manager psqlManager
}

func NewSegmentStorage(m psqlManager) *segmentStorage {
	return &segmentStorage{
		manager: m,
	}
}

func (s segmentStorage) AddSegment(ctx context.Context, dto *domain.SegmentAddDTO) (*domain.Segment, error) {
	segment, _ := s.GetSegmentByName(ctx, &domain.SegmentName{Name: dto.Name})
	if segment != nil {
		return nil, domain.ErrorSegmentExists
	}
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()

	stmt := fmt.Sprintf("INSERT INTO segments (name) VALUES ('%s') RETURNING id", dto.Name)

	var id int64
	err = db.QueryRowxContext(ctx, stmt).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &domain.Segment{
		Id:   id,
		Name: dto.Name,
	}, nil
}
func (s segmentStorage) UpdateSegment(ctx context.Context, dto *domain.SegmentUpdateDTO) (*domain.Segment, error) {
	segment, _ := s.GetSegmentByName(ctx, &domain.SegmentName{Name: dto.NewName})
	if segment != nil {
		return nil, domain.ErrorThisSegmentNameIsUsed
	}
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()

	stmt := fmt.Sprintf("UPDATE segments SET name = '%s' WHERE id = %d", dto.NewName, dto.Id)
	_, err = db.ExecContext(ctx, stmt)
	if err != nil {
		return nil, domain.ErrorNoSuchSegment
	}
	return &domain.Segment{
		Id:   dto.Id,
		Name: dto.NewName,
	}, nil
}
func (s segmentStorage) DeleteSegment(ctx context.Context, dto *domain.SegmentName) error {
	_, err := s.GetSegmentByName(ctx, dto)
	if err != nil {
		return err
	}
	db, err := s.manager.GetDb()
	if err != nil {
		return domain.ErrorWithDataBase
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt := fmt.Sprintf("DELETE FROM user_in_segment WHERE user_in_segment.segment_id IN (SELECT segments.id FROM segments WHERE segments.name = '%s')", dto.Name)
	_, err = tx.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	stmt = fmt.Sprintf("DELETE FROM history WHERE history.segment_id IN (SELECT segments.id FROM segments WHERE segments.name = '%s')", dto.Name)
	_, err = tx.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	stmt = fmt.Sprintf("DELETE FROM segments WHERE name = '%s'", dto.Name)

	_, err = tx.ExecContext(ctx, stmt)
	if err != nil {
		return domain.ErrorNoSuchSegment
	}
	err = tx.Commit()
	return err
}
func (s segmentStorage) GetAllSegments(ctx context.Context) (*[]domain.Segment, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()
	stmt := "SELECT id, name FROM segments"

	rows, err := db.QueryxContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	var segment domain.Segment
	segments := make([]domain.Segment, 0)
	for rows.Next() {
		err := rows.StructScan(&segment)
		if err != nil {
			return nil, err
		}
		segments = append(segments, segment)
	}
	return &segments, nil
}
func (s segmentStorage) GetSegmentById(ctx context.Context, dto *domain.SegmentId) (*domain.Segment, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()
	stmt := fmt.Sprintf("SELECT id, name FROM segments WHERE id = %d", dto.Id)

	var segment domain.Segment
	err = db.QueryRowxContext(ctx, stmt).StructScan(&segment)
	if err != nil {
		return nil, domain.ErrorNoSuchSegment
	}

	return &segment, nil
}

func (s segmentStorage) GetSegmentByName(ctx context.Context, dto *domain.SegmentName) (*domain.Segment, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()
	stmt := fmt.Sprintf("SELECT id, name FROM segments WHERE name = '%s'", dto.Name)

	var segment domain.Segment
	err = db.QueryRowxContext(ctx, stmt).StructScan(&segment)

	if err != nil {
		return nil, domain.ErrorNoSuchSegment
	}

	return &segment, nil
}

func (s segmentStorage) CheckSegmentsExists(ctx context.Context, dto *domain.SegmentNames) error {
	db, err := s.manager.GetDb()
	if err != nil {
		return domain.ErrorWithDataBase
	}
	defer db.Close()

	segments := arrayStrToStr(dto.Names)

	stmt := fmt.Sprintf("SELECT name FROM segments WHERE name IN%s ORDER BY name DESC", segments)

	rows, err := db.QueryxContext(ctx, stmt)
	if err != nil {
		return fmt.Errorf("%s: %v", domain.ErrorNoSuchSegments, dto.Names)
	}

	var (
		segmentName string
		lenght      = len(dto.Names)
	)

	for rows.Next() {
		err := rows.Scan(&segmentName)
		if err != nil {
			return err
		}

		for i := 0; i < lenght; i++ {
			if segmentName == dto.Names[i] {
				dto.Names[i], dto.Names[lenght-1] = dto.Names[lenght-1], dto.Names[i]
				lenght--
				break
			}
		}
	}

	if lenght > 0 {
		return fmt.Errorf("%s: %s", domain.ErrorNoSuchSegments, arrayStrToStr(dto.Names[:lenght]))
	}

	return nil
}

func (s segmentStorage) GetSegmentsIds(ctx context.Context, dto *domain.SegmentNames) (*domain.SegmentIds, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()

	var (
		id         int64
		segmentIds = make([]int64, 0, len(dto.Names))
	)

	stmt := fmt.Sprintf("SELECT id FROM segments WHERE name IN%s", arrayStrToStr(dto.Names))
	rows, err := db.QueryxContext(ctx, stmt)
	if err != nil {
		return nil, domain.ErrorNoSuchSegments
	}

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		segmentIds = append(segmentIds, id)
	}
	return &domain.SegmentIds{
		Ids: segmentIds,
	}, nil
}
