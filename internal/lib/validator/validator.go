package validator

import (
	"fmt"
	"segment-service/internal/core/domain"
)

func ValidateId(id int64) error {
	if id <= 0 {
		return domain.ErrorInvalidId
	}
	return nil
}

func ValidateIds(ids []int64) error {
	for _, id := range ids {
		if id <= 0 {
			return fmt.Errorf("%s: %d", domain.ErrorInvalidId, id)
		}
	}
	return nil
}

func ValidateSegmentName(name string) error {
	if len(name) == 0 {
		return domain.ErrorInvalidName
	}
	return nil
}

func ValidateSegmentNames(names []string) error {
	for _, name := range names {
		if len(name) == 0 {
			return fmt.Errorf("%s: %s", domain.ErrorInvalidName, name)
		}
	}
	return nil
}

func ValidateLimitOffset(limit, offset int) error {
	if limit < 0 {
		return fmt.Errorf("%s: %d", domain.ErrorInvalidLimitValue, limit)
	}
	if offset < 0 {
		return fmt.Errorf("%s: %d", domain.ErrorInvalidOffsetValue, offset)
	}
	return nil
}

func ValidatePercent(percent float32) error {
	if percent <= 0 || percent > 100 {
		return domain.ErrorInvalidpercentValue
	}
	return nil
}

func ValidateHistoryOfUserGetDTO(dto *domain.HistoryOfUserGetDTO) error {
	err := ValidateId(dto.UserId)
	if err != nil {
		return err
	}
	if dto.Month < 0 || dto.Month > 12 {
		fmt.Println(dto.Month)
		return fmt.Errorf("%s: %d", domain.ErrorInvalidMonthValue, dto.Month)
	}
	if dto.Year < 2023 {
		return fmt.Errorf("%s: %d", domain.ErrorInvalidYearValue, dto.Year)
	}
	return nil
}
