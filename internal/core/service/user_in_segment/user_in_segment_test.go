package user_in_segment

import (
	"context"
	"fmt"
	"segment-service/internal/core/domain"
	"segment-service/pkg/config"
	"time"

	history_service "segment-service/internal/core/service/history"
	segment_service "segment-service/internal/core/service/segment"
	user_service "segment-service/internal/core/service/user"
	mock_storage "segment-service/pkg/mocks/storage"

	"testing"

	"github.com/golang/mock/gomock"
)

type storages struct {
	HStorage   *mock_storage.MockHistoryStorage
	UStorage   *mock_storage.MockUserStorager
	SStorage   *mock_storage.MockSegmentStorager
	UISStorage *mock_storage.MockUserInSegmentStorager
}

type services struct {
	HService historyWriter
	UService userOperator
	SService segmentChecker
}

func TestAddUserToSegment(t *testing.T) {
	cases := []struct {
		Name      string
		In        *domain.UserToSegmentsAddDTO
		Out       error
		WhaiError bool
		prepare   func(ctx context.Context, s *storages, in *domain.UserToSegmentsAddDTO)
		TimeLimit time.Duration
	}{
		{
			Name:      "no error",
			In:        &domain.UserToSegmentsAddDTO{UserId: 1, SegmentNames: []string{"ONE", "SELL"}},
			Out:       nil,
			WhaiError: false,
			prepare: func(ctx context.Context, s *storages, in *domain.UserToSegmentsAddDTO) {
				s.HStorage.EXPECT().AddHistory(&domain.HistoryAddDTO{
					UserIds:      []int64{in.UserId},
					SegmentNames: in.SegmentNames,
					Action:       "ADD",
				})
				s.SStorage.EXPECT().CheckSegmentsExists(ctx, &domain.SegmentNames{Names: in.SegmentNames}).Return(nil)
				s.UStorage.EXPECT().CheckUsersExist(ctx, &domain.UsersIds{Ids: []int64{in.UserId}}).Return(nil)
				s.UISStorage.EXPECT().AddUserToSegments(ctx, in)
			},
			TimeLimit: time.Second * 10,
		},
		{
			Name:      "bad user id",
			In:        &domain.UserToSegmentsAddDTO{UserId: 0, SegmentNames: []string{"ONE"}},
			Out:       domain.ErrorInvalidId,
			WhaiError: true,
			prepare:   nil,
			TimeLimit: time.Second * 10,
		},
		{
			Name:      "no such user",
			In:        &domain.UserToSegmentsAddDTO{UserId: 2, SegmentNames: []string{"ONE"}},
			Out:       domain.ErrorNoSuchUsers,
			WhaiError: true,
			prepare: func(ctx context.Context, s *storages, in *domain.UserToSegmentsAddDTO) {
				s.UStorage.EXPECT().CheckUsersExist(ctx, &domain.UsersIds{Ids: []int64{in.UserId}}).Return(domain.ErrorNoSuchUsers)
			},
			TimeLimit: time.Second * 10,
		},
		{
			Name:      "time limit error",
			In:        &domain.UserToSegmentsAddDTO{UserId: 0, SegmentNames: []string{"ONE"}},
			Out:       domain.ErrorTimeOut,
			WhaiError: true,
			prepare:   nil,
			TimeLimit: time.Second * 0,
		},
		{
			Name:      "no such segments",
			In:        &domain.UserToSegmentsAddDTO{UserId: 1, SegmentNames: []string{"No"}},
			Out:       domain.ErrorNoSuchSegments,
			WhaiError: true,
			prepare: func(ctx context.Context, s *storages, in *domain.UserToSegmentsAddDTO) {
				s.SStorage.EXPECT().CheckSegmentsExists(ctx, &domain.SegmentNames{Names: in.SegmentNames}).Return(domain.ErrorNoSuchSegments)
				s.UStorage.EXPECT().CheckUsersExist(ctx, &domain.UsersIds{Ids: []int64{in.UserId}}).Return(nil)
			},
			TimeLimit: time.Second * 10,
		},
		{
			Name:      "no such segments 2",
			In:        &domain.UserToSegmentsAddDTO{UserId: 1, SegmentNames: []string{}},
			Out:       domain.ErrorNoSuchSegments,
			WhaiError: true,
			prepare: func(ctx context.Context, s *storages, in *domain.UserToSegmentsAddDTO) {
				s.SStorage.EXPECT().CheckSegmentsExists(ctx, &domain.SegmentNames{Names: in.SegmentNames}).Return(domain.ErrorNoSuchSegments)
				s.UStorage.EXPECT().CheckUsersExist(ctx, &domain.UsersIds{Ids: []int64{in.UserId}}).Return(nil)
			},
			TimeLimit: time.Second * 10,
		},
	}

	cfg := config.GetConfig()

	for _, cc := range cases {
		t.Run(cc.Name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			storages := &storages{
				UStorage:   mock_storage.NewMockUserStorager(ctl),
				SStorage:   mock_storage.NewMockSegmentStorager(ctl),
				HStorage:   mock_storage.NewMockHistoryStorage(ctl),
				UISStorage: mock_storage.NewMockUserInSegmentStorager(ctl),
			}
			ctx, cancel := context.WithTimeout(context.Background(), cc.TimeLimit)
			defer cancel()
			if cc.prepare != nil {
				cc.prepare(ctx, storages, cc.In)
			}
			services := &services{
				UService: user_service.NewUserService(storages.UStorage),
				SService: segment_service.NewSegmentService(storages.SStorage),
				HService: history_service.NewHistoryService(storages.HStorage, cfg),
			}
			userInSegmentsService := NewUserInSegmentsService(storages.UISStorage, services.UService, services.SService, services.HService)
			err := userInSegmentsService.AddUserToSegments(ctx, cc.In)
			if cc.WhaiError {
				if err.Error() != cc.Out.Error() {
					t.Errorf("not equeal errors: whant - %s; have - %s", err, cc.Out)
				}
			}
		})
	}
}

func TestAddUsersToSegment(t *testing.T) {
	cases := []struct {
		Name      string
		In        *domain.UsersToSegmentsAddDTO
		Out       error
		WhaiError bool
		prepare   func(ctx context.Context, s *storages, in *domain.UsersToSegmentsAddDTO)
		TimeLimit time.Duration
	}{
		{
			Name:      "no error",
			In:        &domain.UsersToSegmentsAddDTO{UserIds: []int64{1, 2, 3}, SegmentNames: []string{"ONE", "SELL"}},
			Out:       nil,
			WhaiError: false,
			prepare: func(ctx context.Context, s *storages, in *domain.UsersToSegmentsAddDTO) {
				s.HStorage.EXPECT().AddHistory(&domain.HistoryAddDTO{
					UserIds:      in.UserIds,
					SegmentNames: in.SegmentNames,
					Action:       "ADD",
				})
				s.SStorage.EXPECT().CheckSegmentsExists(ctx, &domain.SegmentNames{Names: in.SegmentNames}).Return(nil)
				s.UStorage.EXPECT().CheckUsersExist(ctx, &domain.UsersIds{Ids: in.UserIds}).Return(nil)
				s.UISStorage.EXPECT().AddUsersToSegments(ctx, in)
			},
			TimeLimit: time.Second * 10,
		},
		{
			Name:      "bad ids",
			In:        &domain.UsersToSegmentsAddDTO{UserIds: []int64{0, -1}, SegmentNames: []string{"ONE", "SELL"}},
			Out:       fmt.Errorf("%s: %d", domain.ErrorInvalidId, 0),
			WhaiError: true,
			prepare:   nil,
			TimeLimit: time.Second * 10,
		},
		{
			Name:      "no such users",
			In:        &domain.UsersToSegmentsAddDTO{UserIds: []int64{1234}, SegmentNames: []string{"ONE", "SELL"}},
			Out:       domain.ErrorNoSuchUsers,
			WhaiError: true,
			prepare: func(ctx context.Context, s *storages, in *domain.UsersToSegmentsAddDTO) {
				s.UStorage.EXPECT().CheckUsersExist(ctx, &domain.UsersIds{Ids: in.UserIds}).Return(domain.ErrorNoSuchUsers)
			},
			TimeLimit: time.Second * 10,
		},
		{
			Name:      "no such users 2",
			In:        &domain.UsersToSegmentsAddDTO{UserIds: []int64{}, SegmentNames: []string{"ONE", "SELL"}},
			Out:       domain.ErrorNoSuchUsers,
			WhaiError: true,
			prepare: func(ctx context.Context, s *storages, in *domain.UsersToSegmentsAddDTO) {
				s.UStorage.EXPECT().CheckUsersExist(ctx, &domain.UsersIds{Ids: in.UserIds}).Return(domain.ErrorNoSuchUsers)
			},
			TimeLimit: time.Second * 10,
		},
		{
			Name:      "no such segments",
			In:        &domain.UsersToSegmentsAddDTO{UserIds: []int64{1, 2}, SegmentNames: []string{"No"}},
			Out:       domain.ErrorNoSuchSegments,
			WhaiError: true,
			prepare: func(ctx context.Context, s *storages, in *domain.UsersToSegmentsAddDTO) {
				s.UStorage.EXPECT().CheckUsersExist(ctx, &domain.UsersIds{Ids: in.UserIds}).Return(nil)
				s.SStorage.EXPECT().CheckSegmentsExists(ctx, &domain.SegmentNames{Names: in.SegmentNames}).Return(domain.ErrorNoSuchSegments)
			},
			TimeLimit: time.Second * 10,
		},
		{
			Name:      "time out",
			In:        &domain.UsersToSegmentsAddDTO{UserIds: []int64{1, 2}, SegmentNames: []string{"No"}},
			Out:       domain.ErrorTimeOut,
			WhaiError: true,
			prepare:   nil,
			TimeLimit: 0,
		},
	}

	cfg := config.GetConfig()

	for _, cc := range cases {
		t.Run(cc.Name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			storages := &storages{
				UStorage:   mock_storage.NewMockUserStorager(ctl),
				SStorage:   mock_storage.NewMockSegmentStorager(ctl),
				HStorage:   mock_storage.NewMockHistoryStorage(ctl),
				UISStorage: mock_storage.NewMockUserInSegmentStorager(ctl),
			}
			ctx, cancel := context.WithTimeout(context.Background(), cc.TimeLimit)
			defer cancel()
			if cc.prepare != nil {
				cc.prepare(ctx, storages, cc.In)
			}
			services := &services{
				UService: user_service.NewUserService(storages.UStorage),
				SService: segment_service.NewSegmentService(storages.SStorage),
				HService: history_service.NewHistoryService(storages.HStorage, cfg),
			}
			userInSegmentsService := NewUserInSegmentsService(storages.UISStorage, services.UService, services.SService, services.HService)
			err := userInSegmentsService.AddUsersToSegments(ctx, cc.In)
			if cc.WhaiError {
				if err.Error() != cc.Out.Error() {
					t.Errorf("not equeal errors: whant - %s; have - %s", cc.Out, err)
				}
			}
		})
	}
}
