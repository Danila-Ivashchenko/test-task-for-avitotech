package domain

import "errors"

var (
	ErrorTimeOut      = errors.New("time out error")
	ErrorWithDataBase = errors.New("error with data base")
	ErrorUserExists   = errors.New("user exists")
	ErrorNoUsers      = errors.New("no users")

	ErrorNoSuchUser = errors.New("no such user")

	ErrorNoSuchSegment         = errors.New("no such segment")
	ErrorSegmentExists         = errors.New("segment with this name already exists")
	ErrorThisSegmentNameIsUsed = errors.New("this segment name is used")
	ErrorNoSuchSegments        = errors.New("no such segments")

	ErrorNoSuchUserOrSegments   = errors.New("no such user or segments")
	ErrorUserNotInSegments      = errors.New("user is not in some segments")
	ErrorSomeUsersAlreadyInSegments = errors.New("some users already in segments")

	ErrorNoHistoryResult = errors.New("this user has not records in this month")
)
