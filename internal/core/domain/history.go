package domain

import (
	"time"
)

type HistoryResponce struct {
	URL string `json:"url"`
}

type HistoryOfUser struct {
	SegmentName string    `db:"name"`
	Action      string    `db:"action"`
	Time        time.Time `db:"action_time"`
}

func (h HistoryOfUser) Values() []string {
	return []string{h.SegmentName, h.Action, h.Time.Format("2006-01-02 15:04:05")}
}

type HistoryAddDTO struct {
	UserIds      []int64
	SegmentNames []string
	Action       string
}

type HistoryOfUserGetDTO struct {
	UserId int64 `json:"user_id"`
	Month  int   `json:"month"`
	Year   int   `json:"year"`
}
