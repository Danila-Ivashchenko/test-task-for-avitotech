package domain

import "errors"

var (
	ErrorTimeOut      = errors.New("time out error")
	ErrorWithDataBase = errors.New("error with data base")
	ErrorUserExists   = errors.New("user exists")
	ErrorInvalidId    = errors.New("invalid id")
	ErrorNoUsers      = errors.New("no users")

	ErrorNoSuchUser = errors.New("no such user")
	ErrorNoSuchUsers = errors.New("no such users")

	ErrorNoSuchSegment         = errors.New("no such segment")
	ErrorInvalidName           = errors.New("invalid name")
	ErrorSegmentExists         = errors.New("segment with this name already exists")
	ErrorThisSegmentNameIsUsed = errors.New("this segment name is used")
	ErrorNoSuchSegments        = errors.New("no such segments")

	ErrorNoSuchUserOrSegments       = errors.New("no such user or segments")
	ErrorUserNotInSegments          = errors.New("user is not in some segments")
	ErrorSomeUsersAlreadyInSegments = errors.New("some users already in segments")

	ErrorNoHistoryResult = errors.New("this user has not records in this month")

	ErrorInvalidLimitValue   = errors.New("invalid value of limit")
	ErrorInvalidOffsetValue  = errors.New("invalid value of offset")
	ErrorInvalidpercentValue = errors.New("invalid value of percent")
	ErrorInvalidMonthValue   = errors.New("invalid value of month")
	ErrorInvalidYearValue    = errors.New("invalid value of year")
)
