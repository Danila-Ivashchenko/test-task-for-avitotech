package domain

type Segment struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type SegmentAddDTO struct {
	Name string `json:"name" db:"name"`
}

type SegmentUpdateDTO struct {
	Id      int64  `json:"id"`
	NewName string `json:"new_name"`
}

type SegmentId struct {
	Id int64 `json:"id"`
}

type SegmentIds struct {
	Ids []int64 `json:"id"`
}

type SegmentName struct {
	Name string `json:"name"`
}

type SegmentNames struct {
	Names []string `json:"segments"`
}
