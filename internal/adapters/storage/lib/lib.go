package lib

import (
	"fmt"
	"segment-service/internal/core/domain"

	"github.com/jmoiron/sqlx"
)

type PsqlManager interface {
	GetDb() (*sqlx.DB, error)
}

func ArrayInt64ToStr(arr []int64) string {
	result := "("

	for i := range arr {
		if i == 0 {
			result += fmt.Sprintf("%d", arr[i])
		} else {
			result += fmt.Sprintf(", %d", arr[i])
		}
	}
	return result + ")"
}

func ArrayStrToStr(arr []string) string {
	result := "("

	for i := range arr {
		if i == 0 {
			result += fmt.Sprintf("'%s'", arr[i])
		} else {
			result += fmt.Sprintf(", '%s'", arr[i])
		}
	}
	return result + ")"
}

func MakeUserInSegmentPairs(userId int64, segmentIds []int64) string {
	result := ""
	for i := range segmentIds {
		if i == 0 {
			result += fmt.Sprintf("(%d, %d) ", userId, segmentIds[i])
		} else {
			result += fmt.Sprintf(", (%d, %d)", userId, segmentIds[i])
		}
	}
	return result
}

func MergeUsersAndSegments(users, segments []int64) string {
	result := ""
	for i, user := range users {
		if i > 0 {
			result += ", "
		}
		for j, segment := range segments {
			if j == 0 {
				result += fmt.Sprintf("(%d, %d)", user, segment)
			} else {
				result += fmt.Sprintf(", (%d, %d)", user, segment)
			}
		}
	}
	return result
}

func MakeHistoryRecords(users, segments []int64, action string) string {
	result := ""
	for i, user := range users {
		if i > 0 {
			result += ", "
		}
		for j, segment := range segments {
			if j == 0 {
				result += fmt.Sprintf("(%d, %d, '%s')", user, segment, action)
			} else {
				result += fmt.Sprintf(", (%d, %d, '%s')", user, segment, action)
			}
		}
	}

	return result
}

func GetIdsFromRows(rows *sqlx.Rows, size, cap int) (*domain.UsersIds, error) {
	var id int64
	ids := make([]int64, size, cap)
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return &domain.UsersIds{
		Ids: ids,
	}, nil
}
