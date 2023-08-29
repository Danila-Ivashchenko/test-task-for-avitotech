package storage

import (
	"context"
	"fmt"
	"segment-service/internal/core/domain"
)

type userInSegmentStorage struct {
	manager psqlManager
}

func NewUserInSegmentStorage(m psqlManager) *userInSegmentStorage {
	return &userInSegmentStorage{
		manager: m,
	}
}
func (s userInSegmentStorage) getSegmentIds(ctx context.Context, dto *domain.SegmentNames) ([]int64, error) {
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

func (s userInSegmentStorage) AddUserToSegments(ctx context.Context, dto *domain.UserToSegmentsAddDTO) error {
	db, err := s.manager.GetDb()
	if err != nil {
		return domain.ErrorWithDataBase
	}
	defer db.Close()

	segmentIds, err := s.getSegmentIds(ctx, &domain.SegmentNames{Names: dto.SegmentNames})
	if err != nil {
		return err
	}

	stmt := fmt.Sprintf("INSERT INTO user_in_segment (user_id, segment_id) VALUES %s", makeUserInSegmentPairs(dto.UserId, segmentIds))

	_, err = db.ExecContext(ctx, stmt)
	if err != nil {
		return domain.ErrorSomeUsersAlreadyInSegments
	}

	return nil
}

func (s userInSegmentStorage) AddUsersToSegments(ctx context.Context, dto *domain.UsersToSegmentsAddDTO) error {
	db, err := s.manager.GetDb()
	if err != nil {
		return domain.ErrorWithDataBase
	}
	defer db.Close()

	segmentIds, err := s.getSegmentIds(ctx, &domain.SegmentNames{Names: dto.SegmentNames})
	if err != nil {
		return err
	}

	stmt := fmt.Sprintf("INSERT INTO user_in_segment (user_id, segment_id) VALUES %s", mergeUsersAndSegments(dto.UserIds, segmentIds))

	_, err = db.ExecContext(ctx, stmt)
	if err != nil {
		return domain.ErrorSomeUsersAlreadyInSegments
	}

	return nil
}
func (s userInSegmentStorage) checkUserInSegments(ctx context.Context, dto *domain.UserInSegmentsCheckDTO) error {
	db, err := s.manager.GetDb()
	if err != nil {
		return domain.ErrorWithDataBase
	}
	defer db.Close()

	stmt := fmt.Sprintf(`SELECT COUNT(*) FROM user_in_segment WHERE user_id = %d AND segment_id IN%s`, dto.UserId, arrayInt64ToStr(dto.SegmentIds))
	var (
		idsLen int64
	)
	err = db.QueryRow(stmt).Scan(&idsLen)
	if err != nil {
		return err
	}
	if len(dto.SegmentIds) != int(idsLen) {
		return domain.ErrorUserNotInSegments
	}
	return nil
}

func (s userInSegmentStorage) DeleteUserFromSegments(ctx context.Context, dto *domain.UserFromSegmentsDeleteDTO) error {
	db, err := s.manager.GetDb()
	if err != nil {
		return domain.ErrorWithDataBase
	}
	defer db.Close()

	segmentIds, err := s.getSegmentIds(ctx, &domain.SegmentNames{Names: dto.SegmentNames})
	if err != nil {
		return err
	}
	err = s.checkUserInSegments(ctx, &domain.UserInSegmentsCheckDTO{UserId: dto.UserId, SegmentIds: segmentIds})
	if err != nil {
		return err
	}

	stmt := fmt.Sprintf("DELETE FROM user_in_segment WHERE user_id = %d AND segment_id IN%s", dto.UserId, arrayInt64ToStr(segmentIds))

	_, err = db.ExecContext(ctx, stmt)
	if err != nil {
		return domain.ErrorNoSuchUserOrSegments
	}

	return nil
}
func (s userInSegmentStorage) GetUserInSegments(ctx context.Context, dto *domain.UserId) (*domain.UserInSegments, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()

	stmt := fmt.Sprintf("SELECT segments.name FROM user_in_segment, segments WHERE user_in_segment.user_id = %d AND user_in_segment.segment_id = segments.id", dto.Id)
	rows, err := db.QueryxContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	var (
		segmentName   string
		segmentsNames = make([]string, 0)
	)

	for rows.Next() {
		err := rows.Scan(&segmentName)
		if err != nil {
			return nil, err
		}
		segmentsNames = append(segmentsNames, segmentName)
	}
	return &domain.UserInSegments{
		UserId:       dto.Id,
		SegmentNames: segmentsNames,
	}, nil
}

func (s userInSegmentStorage) GetUsersInSegment(ctx context.Context, dto *domain.SegmentName) (*domain.UsersInSegment, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()

	stmt := fmt.Sprintf("SELECT user_id FROM user_in_segment, segments WHERE user_in_segment.segment_id = segments.id AND segments.name = '%s'", dto.Name)
	rows, err := db.QueryxContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	var (
		userId   int64
		usersIds = make([]int64, 0)
	)

	for rows.Next() {
		err := rows.Scan(&userId)
		if err != nil {
			return nil, err
		}
		usersIds = append(usersIds, userId)
	}
	return &domain.UsersInSegment{
		SegmentName: dto.Name,
		UsersIds:    usersIds,
	}, nil
}

func (s userInSegmentStorage) AddUsersWithLimitOffsetToSegments(ctx context.Context, dto *domain.UsersWithLimitOffsetToSegments) (*domain.UsersIds, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()

	segmentIds, err := s.getSegmentIds(ctx, &domain.SegmentNames{Names: dto.SegmentNames})
	if err != nil {
		return nil, err
	}
	
	additionCase := fmt.Sprintf(" WHERE id NOT IN(SELECT user_id FROM user_in_segment WHERE segment_id IN%s)", arrayInt64ToStr(segmentIds))
	if dto.Rand {
		additionCase += " ORDER BY RANDOM()"
	}
	if dto.Limit > 0 {
		additionCase += fmt.Sprintf(` LIMIT %d`, dto.Limit)
	}
	if dto.Offset > 0 {
		additionCase += fmt.Sprintf(` OFFSET %d`, dto.Offset)
	}

	users, err := s.getUsersWithAdditionCase(ctx, additionCase, dto.Limit)
	
	if err != nil {
		return nil, err
	}
	if len(users.Ids) == 0 {
		return nil, domain.ErrorNoUsers
	}
	fmt.Println(users.Ids)
	stmt := fmt.Sprintf("INSERT INTO user_in_segment (user_id, segment_id) VALUES %s", mergeUsersAndSegments(users.Ids, segmentIds))

	_, err = db.ExecContext(ctx, stmt)
	if err != nil {
		return nil, domain.ErrorSomeUsersAlreadyInSegments
	}

	return users, nil
}

func (s userInSegmentStorage) getUsersWithAdditionCase(ctx context.Context, additionCase string, cap int) (*domain.UsersIds, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()
	stmt := "SELECT id FROM users " + additionCase
	fmt.Println(stmt)
	rows, err := db.QueryxContext(ctx, stmt)
	if err != nil {
		return nil, err
	}

	return getIdsFromRows(rows, 0, cap)
}

func (s userInSegmentStorage) AddPercentOfUsersToSegments(ctx context.Context, dto *domain.PercentOfUsersToSegmentsDTO) (*domain.UsersIds, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()

	segmentIds, err := s.getSegmentIds(ctx, &domain.SegmentNames{Names: dto.SegmentNames})
	if err != nil {
		return nil, err
	}

	additionCase := fmt.Sprintf("WHERE %s ORDER BY RANDOM() LIMIT (SELECT ROUND(COUNT(*) * %f / 100) FROM users)",
		fmt.Sprintf(" id NOT IN(SELECT user_id FROM user_in_segment WHERE segment_id IN%s)", arrayInt64ToStr(segmentIds)),
		dto.Percent,
	)

	users, err := s.getUsersWithAdditionCase(ctx, additionCase, 0)
	if err != nil {
		return nil, err
	}
	if len(users.Ids) == 0 {
		return nil, domain.ErrorNoUsers
	}

	stmt := fmt.Sprintf("INSERT INTO user_in_segment (user_id, segment_id) VALUES %s", mergeUsersAndSegments(users.Ids, segmentIds))

	_, err = db.ExecContext(ctx, stmt)
	if err != nil {
		return nil, domain.ErrorSomeUsersAlreadyInSegments
	}

	return users, nil
}
