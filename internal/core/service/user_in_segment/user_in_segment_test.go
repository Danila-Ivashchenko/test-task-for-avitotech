package user_in_segment

import (
	"context"
	"segment-service/internal/core/domain"
	history_service "segment-service/internal/core/service/history"
	segment_service "segment-service/internal/core/service/segment"
	user_service "segment-service/internal/core/service/user"
	"segment-service/pkg/config"
	"time"

	mock_storage "segment-service/internal/mocks/storage"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAddUserToSegment(t *testing.T) {
	cases := []struct {
		Name        string
		In          *domain.UserToSegmentsAddDTO
		Out         error
		WhaiError   bool
		ServicesUse int
		TimeLimit   time.Duration
	}{
		{
			Name:        "no error",
			In:          &domain.UserToSegmentsAddDTO{UserId: 1, SegmentNames: []string{"ONE", "SELL"}},
			Out:         nil,
			WhaiError:   false,
			ServicesUse: 4,
			TimeLimit:   time.Second * 10,
		},
		{
			Name:        "bad user id",
			In:          &domain.UserToSegmentsAddDTO{UserId: 0, SegmentNames: []string{"ONE"}},
			Out:         domain.ErrorInvalidId,
			WhaiError:   true,
			ServicesUse: 0,
			TimeLimit:   time.Second * 10,
		},
		{
			Name:        "time limit error",
			In:          &domain.UserToSegmentsAddDTO{UserId: 0, SegmentNames: []string{"ONE"}},
			Out:         domain.ErrorTimeOut,
			WhaiError:   true,
			ServicesUse: 0,
			TimeLimit:   time.Second * 0,
		},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	cfg := config.GetConfig()
	userStorage := mock_storage.NewMockUserStorager(ctl)
	userService := user_service.NewUserService(userStorage)
	segmentStorage := mock_storage.NewMockSegmentStorager(ctl)
	segmentService := segment_service.NewSegmentService(segmentStorage)
	historyStorage := mock_storage.NewMockHistoryStorage(ctl)
	historyService := history_service.NewHistoryService(historyStorage, cfg)
	userInSegmentsStorage := mock_storage.NewMockUserInSegmentStorager(ctl)

	userInSegmentsService := NewUserInSegmentsService(userInSegmentsStorage, userService, segmentService, historyService)

	for i := 0; i < len(cases); i++ {
		t.Run(cases[i].Name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), cases[i].TimeLimit)
			defer cancel()

			if cases[i].ServicesUse >= 1 {
				userStorage.EXPECT().CheckUserExists(ctx, &domain.UsersIds{Ids: []int64{cases[i].In.UserId}})
			}
			if cases[i].ServicesUse >= 2 {
				segmentStorage.EXPECT().CheckSegmentsExists(ctx, &domain.SegmentNames{Names: cases[i].In.SegmentNames})
			}
			if cases[i].ServicesUse >= 3 {
				userInSegmentsStorage.EXPECT().AddUserToSegments(ctx, cases[i].In)
			}
			if cases[i].ServicesUse >= 4 {
				historyStorage.EXPECT().AddHistory(&domain.HistoryAddDTO{
					UserIds:      []int64{cases[i].In.UserId},
					SegmentNames: cases[i].In.SegmentNames,
					Action:       "ADD",
				})
			}

			err := userInSegmentsService.AddUserToSegments(ctx, cases[i].In)
			if cases[i].WhaiError {
				if err.Error() != cases[i].Out.Error() {
					t.Errorf("not equeal errors: whant - %s; have - %s", err, cases[i].Out)
				}
			}
		})
	}
}
