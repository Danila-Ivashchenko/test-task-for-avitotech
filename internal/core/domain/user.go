package domain

import "time"

type User struct {
	Id        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UserId struct {
	Id int64 `json:"id" db:"id"`
}

type UsersIds struct {
	Ids []int64 `json:"ids" db:"id"`
}

type UsersGetPercentDTO struct {
	Percent float32 `json:"persent"`
}

type UserAffected struct {
	Affected int64 `json:"affected"`
	Ignored int64 `json:"ignored"`
}

type LinitOffset struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Rand bool `json:"random"`
}

