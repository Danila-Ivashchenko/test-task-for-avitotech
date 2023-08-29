package storage

import (
	"fmt"
	"segment-service/internal/core/domain"

	"github.com/jmoiron/sqlx"
)

type psqlManager interface {
	GetDb() (*sqlx.DB, error)
}

func arrayInt64ToStr(arr []int64) string {
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

func arrayStrToStr(arr []string) string {
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

func makeUserInSegmentPairs(userId int64, segmentIds []int64) string {
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

func mergeUsersAndSegments(users, segments []int64) string {
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

func makeHistoryRecords(users, segments []int64, action string) string {
	result := ""
	for i, user := range users {
		if i > 0 {
			result += ", "
		}
		for j, segment := range segments {
			if j == 0 {
				result = fmt.Sprintf("(%d, %d, '%s')", user, segment, action)
			} else {
				result = fmt.Sprintf(", (%d, %d, '%s')", user, segment, action)
			}
		}
	}

	return result
}

func getIdsFromRows(rows *sqlx.Rows, size, cap int) (*domain.UsersIds, error) {
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
