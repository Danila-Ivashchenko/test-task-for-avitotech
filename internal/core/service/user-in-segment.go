package service

import (
	"context"
	"segment-service/internal/core/domain"
	"segment-service/internal/core/ports/storage"
)

type userOperator interface {
	CheckUserExists(context.Context, *domain.UsersIds) (*domain.UsersIds, error)
}

type segmentChecker interface {
	CheckSegmentsExists(ctx context.Context, dto *domain.SegmentNames) error
}

type historyWriter interface {
	AddHistory(dto *domain.HistoryAddDTO)
}

type userInSegmentsService struct {
	storage        storage.UserInSegmentStorager
	userOperator   userOperator
	segmentChecker segmentChecker
	historyWriter  historyWriter
}

func NewUserInSegmentsService(s storage.UserInSegmentStorager, uChecker userOperator, sChecker segmentChecker, hWriter historyWriter) *userInSegmentsService {
	return &userInSegmentsService{
		storage:        s,
		userOperator:   uChecker,
		segmentChecker: sChecker,
		historyWriter:  hWriter,
	}
}

func (s userInSegmentsService) AddUserToSegments(ctx context.Context, dto *domain.UserToSegmentsAddDTO) error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)

		_, err := s.userOperator.CheckUserExists(ctx, &domain.UsersIds{Ids: []int64{dto.UserId}})
		if err != nil {
			errCh <- domain.ErrorNoSuchUser
			return
		}

		err = s.segmentChecker.CheckSegmentsExists(ctx, &domain.SegmentNames{Names: dto.SegmentNames})
		if err != nil {
			errCh <- err
			return
		}

		err = s.storage.AddUserToSegments(ctx, dto)
		if err != nil {
			errCh <- err
			return
		}
		go s.historyWriter.AddHistory(
			&domain.HistoryAddDTO{
				UserIds:      []int64{dto.UserId},
				SegmentNames: dto.SegmentNames,
				Action:       "ADD",
			},
		)
		errCh <- nil
	}()

	select {
	case <-ctx.Done():
		return domain.ErrorTimeOut
	case err := <-errCh:
		return err
	}
}

func (s userInSegmentsService) AddUsersToSegments(ctx context.Context, dto *domain.UsersToSegmentsAddDTO) error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)

		_, err := s.userOperator.CheckUserExists(ctx, &domain.UsersIds{Ids: dto.UserIds})
		if err != nil {
			errCh <- domain.ErrorNoSuchUser
			return
		}

		err = s.segmentChecker.CheckSegmentsExists(ctx, &domain.SegmentNames{Names: dto.SegmentNames})
		if err != nil {
			errCh <- domain.ErrorNoSuchUser
			return
		}

		err = s.storage.AddUsersToSegments(ctx, dto)
		errCh <- err
		if err != nil {
			go s.historyWriter.AddHistory(
				&domain.HistoryAddDTO{
					UserIds:      dto.UserIds,
					SegmentNames: dto.SegmentNames,
					Action:       "ADD",
				},
			)
		}
	}()

	select {
	case <-ctx.Done():
		return domain.ErrorTimeOut
	case err := <-errCh:
		return err
	}
}

func (s userInSegmentsService) AddPersentOfUsersToSegments(ctx context.Context, dto *domain.PersentOfUsersToSegmentsDTO) error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)

		err := s.segmentChecker.CheckSegmentsExists(ctx, &domain.SegmentNames{Names: dto.SegmentNames})
		if err != nil {
			errCh <- domain.ErrorNoSuchUser
			return
		}

		users, err := s.storage.AddPersentOfUsersToSegments(ctx, dto)
		errCh <- err
		if err == nil {
			go s.historyWriter.AddHistory(
				&domain.HistoryAddDTO{
					UserIds:      users.Ids,
					SegmentNames: dto.SegmentNames,
					Action:       "ADD",
				},
			)
		}
	}()

	select {
	case <-ctx.Done():
		return domain.ErrorTimeOut
	case err := <-errCh:
		return err
	}
}

func (s userInSegmentsService) AddUsersWithLimitOffsetToSegments(ctx context.Context, dto *domain.UsersWithLimitOffsetToSegments) error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)

		err := s.segmentChecker.CheckSegmentsExists(ctx, &domain.SegmentNames{Names: dto.SegmentNames})
		if err != nil {
			errCh <- err
			return
		}

		users, err := s.storage.AddUsersWithLimitOffsetToSegments(ctx, dto)
		errCh <- err
		if err == nil {

			go s.historyWriter.AddHistory(
				&domain.HistoryAddDTO{
					UserIds:      users.Ids,
					SegmentNames: dto.SegmentNames,
					Action:       "ADD",
				},
			)
		}
	}()

	select {
	case <-ctx.Done():
		return domain.ErrorTimeOut
	case err := <-errCh:
		return err
	}
}

func (s userInSegmentsService) DeleteUserFromSegments(ctx context.Context, dto *domain.UserFromSegmentsDeleteDTO) error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)

		_, err := s.userOperator.CheckUserExists(ctx, &domain.UsersIds{Ids: []int64{dto.UserId}})
		if err != nil {
			errCh <- domain.ErrorNoSuchUser
			return
		}

		err = s.segmentChecker.CheckSegmentsExists(ctx, &domain.SegmentNames{Names: dto.SegmentNames})
		if err != nil {
			errCh <- err
			return
		}
		err = s.storage.DeleteUserFromSegments(ctx, dto)
		if err == nil {
			go s.historyWriter.AddHistory(
				&domain.HistoryAddDTO{
					UserIds:      []int64{dto.UserId},
					SegmentNames: dto.SegmentNames,
					Action:       "DEL",
				},
			)
		}
		errCh <- err
	}()

	select {
	case <-ctx.Done():
		return domain.ErrorTimeOut
	case err := <-errCh:
		return err
	}
}

func (s userInSegmentsService) GetUserInSegments(ctx context.Context, dto *domain.UserId) (*domain.UserInSegments, error) {
	resultCh := make(chan *domain.UserInSegments)
	errCh := make(chan error)

	go func() {
		defer close(resultCh)
		defer close(errCh)

		_, err := s.userOperator.CheckUserExists(ctx, &domain.UsersIds{Ids: []int64{dto.Id}})
		if err != nil {
			errCh <- domain.ErrorNoSuchUser
			return
		}

		result, err := s.storage.GetUserInSegments(ctx, dto)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- result
	}()

	select {
	case <-ctx.Done():
		return nil, domain.ErrorTimeOut
	case err := <-errCh:
		return nil, err
	case result := <-resultCh:
		return result, nil
	}
}

func (s userInSegmentsService) GetUsersInSegment(ctx context.Context, dto *domain.SegmentName) (*domain.UsersInSegment, error) {
	resultCh := make(chan *domain.UsersInSegment)
	errCh := make(chan error)

	go func() {
		defer close(resultCh)
		defer close(errCh)

		err := s.segmentChecker.CheckSegmentsExists(ctx, &domain.SegmentNames{Names: []string{dto.Name}})
		if err != nil {
			errCh <- domain.ErrorNoSuchSegment
			return
		}

		result, err := s.storage.GetUsersInSegment(ctx, dto)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- result
	}()

	select {
	case <-ctx.Done():
		return nil, domain.ErrorTimeOut
	case err := <-errCh:
		return nil, err
	case result := <-resultCh:
		return result, nil
	}
}