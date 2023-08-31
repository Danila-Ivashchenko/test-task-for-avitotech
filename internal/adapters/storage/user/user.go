package user

import (
	"context"
	"fmt"
	"segment-service/internal/adapters/storage/lib"
	"segment-service/internal/core/domain"
)

type userStorage struct {
	manager lib.PsqlManager
}

func NewUserStorage(m lib.PsqlManager) *userStorage {
	return &userStorage{
		manager: m,
	}
}

func (s userStorage) AddUsers(ctx context.Context, dto *domain.UsersIds) (*domain.UserAffected, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()
	stmt := "INSERT INTO users (id) VALUES " + arrayToValues(dto.Ids)
	stmt += " ON CONFLICT (id) DO NOTHING"

	result, err := db.ExecContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	affected, err := result.RowsAffected()
	ignored := len(dto.Ids) - int(affected)
	return &domain.UserAffected{
		Affected: affected,
		Ignored:  int64(ignored),
	}, err
}

func (s userStorage) GetUser(ctx context.Context, dto *domain.UserId) (*domain.User, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()
	stmt := fmt.Sprintf(`SELECT id, created_at FROM users WHERE id = %d`, dto.Id)

	user := domain.User{}
	err = db.QueryRowxContext(ctx, stmt).StructScan(&user)
	if err != nil {
		return nil, domain.ErrorNoSuchUser
	}
	return &user, nil
}

func (s userStorage) GetUsersIds(ctx context.Context, dto *domain.LinitOffset) (*domain.UsersIds, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()
	stmt := "SELECT id FROM users"
	if dto.Rand {
		stmt += " ORDER BY RANDOM()"
	}
	if dto.Limit > 0 {
		stmt += fmt.Sprintf(` LIMIT %d`, dto.Limit)
	}
	if dto.Offset > 0 {
		stmt += fmt.Sprintf(` OFFSET %d`, dto.Offset)
	}

	rows, err := db.QueryxContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	return lib.GetIdsFromRows(rows, 0, dto.Limit)
}
func (s userStorage) GetPercentOfUsersIds(ctx context.Context, dto *domain.UsersGetPercentDTO) (*domain.UsersIds, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()
	stmt := fmt.Sprintf("SELECT id FROM users ORDER BY RANDOM() LIMIT (SELECT ROUND(COUNT(*) * %f / 100) FROM users)", dto.Percent)

	rows, err := db.QueryxContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	return lib.GetIdsFromRows(rows, 0, 0)

}

func (s userStorage) DeleteUsers(ctx context.Context, dto *domain.UsersIds) (*domain.UserAffected, error) {
	db, err := s.manager.GetDb()
	if err != nil {
		return nil, domain.ErrorWithDataBase
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	stmt := "DELETE FROM history WHERE user_id IN" + lib.ArrayInt64ToStr(dto.Ids)
	tx.ExecContext(ctx, stmt)

	stmt = "DELETE FROM user_in_segment WHERE user_id IN" + lib.ArrayInt64ToStr(dto.Ids)
	tx.ExecContext(ctx, stmt)

	stmt = "DELETE FROM users WHERE id IN" + lib.ArrayInt64ToStr(dto.Ids)
	result, err := tx.ExecContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	affected, err := result.RowsAffected()
	ignored := len(dto.Ids) - int(affected)
	return &domain.UserAffected{
		Affected: affected,
		Ignored:  int64(ignored),
	}, err
}

func (s userStorage) CheckUsersExist(ctx context.Context, dto *domain.UsersIds) error {
	db, err := s.manager.GetDb()
	if err != nil {
		return domain.ErrorWithDataBase
	}
	defer db.Close()

	stmt := "SELECT id FROM users WHERE id IN" + lib.ArrayInt64ToStr(dto.Ids)

	rows, err := db.QueryxContext(ctx, stmt)
	if err != nil {
		return domain.ErrorNoSuchUsers
	}
	users, err := lib.GetIdsFromRows(rows, 0, len(dto.Ids))
	if len(users.Ids) != len(dto.Ids) {
		return domain.ErrorNoSuchUsers
	}
	return nil
}

func arrayToValues(arr []int64) string {
	result := ""
	for i := range arr {
		if i == 0 {
			result += fmt.Sprintf("(%d)", arr[i])
		} else {
			result += fmt.Sprintf(", (%d)", arr[i])
		}
	}
	return result
}
