package domain

type UserInSegments struct {
	UserId       int64    `json:"user_id"`
	SegmentNames []string `json:"segment_names"`
}

type UsersInSegment struct {
	SegmentName string  `json:"segment_name"`
	UsersIds    []int64 `json:"ids"`
}

type UserToSegmentsAddDTO struct {
	UserId       int64    `json:"user_id"`
	SegmentNames []string `json:"segment_names"`
}

type UserInSegmentsCheckDTO struct {
	UserId     int64
	SegmentIds []int64
}

type UsersToSegmentsAddDTO struct {
	UserIds      []int64  `json:"user_ids"`
	SegmentNames []string `json:"segment_names"`
}

type PersentOfUsersToSegmentsDTO struct {
	Percent      float32  `json:"persent"`
	SegmentNames []string `json:"segment_names"`
}

type UsersWithLimitOffsetToSegments struct {
	Limit        int      `json:"limit"`
	Offset       int      `json:"offset"`
	Rand         bool     `json:"random"`
	SegmentNames []string `json:"segment_names"`
}

type UserFromSegmentsDeleteDTO struct {
	UserId       int64    `json:"user_id"`
	SegmentNames []string `json:"segment_names"`
}

type UserInSegmentsGetDTO struct {
	UserId int64 `json:"user_id"`
}
