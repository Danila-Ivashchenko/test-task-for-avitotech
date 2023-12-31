package user

import (
	"context"
	"segment-service/internal/core/domain"
	"segment-service/internal/core/ports/storage"
	"segment-service/internal/lib/validator"
)

type userService struct {
	storage storage.UserStorager
}

func NewUserService(s storage.UserStorager) *userService {
	return &userService{
		storage: s,
	}
}

func (s userService) AddUsers(ctx context.Context, dto *domain.UsersIds) (*domain.UserAffected, error) {
	resultCh := make(chan *domain.UserAffected)
	errCh := make(chan error)

	go func() {
		defer close(resultCh)
		defer close(errCh)

		err := validator.ValidateIds(dto.Ids)
		if err != nil {
			errCh <- err
			return
		}
		result, err := s.storage.AddUsers(ctx, dto)
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

func (s userService) CheckUsersExist(ctx context.Context, dto *domain.UsersIds) error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)
		err := validator.ValidateIds(dto.Ids)
		if err != nil {
			errCh <- err
			return
		}
		errCh <- s.storage.CheckUsersExist(ctx, dto)
	}()

	select {
	case <-ctx.Done():
		return domain.ErrorTimeOut
	case err := <-errCh:
		return err
	}
}

func (s userService) GetUser(ctx context.Context, dto *domain.UserId) (*domain.User, error) {
	resultCh := make(chan *domain.User)
	errCh := make(chan error)

	go func() {
		defer close(resultCh)
		defer close(errCh)
		err := validator.ValidateId(dto.Id)
		if err != nil {
			errCh <- err
			return
		}

		user, err := s.storage.GetUser(ctx, dto)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- user
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

func (s userService) GetUsersIds(ctx context.Context, dto *domain.LinitOffset) (*domain.UsersIds, error) {
	resultCh := make(chan *domain.UsersIds)
	errCh := make(chan error)

	go func() {
		defer close(resultCh)
		defer close(errCh)

		err := validator.ValidateLimitOffset(dto.Limit, dto.Offset)
		if err != nil {
			errCh <- err
			return
		}
		userIds, err := s.storage.GetUsersIds(ctx, dto)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- userIds
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

func (s userService) GetPercentOfUsersIds(ctx context.Context, dto *domain.UsersGetPercentDTO) (*domain.UsersIds, error) {
	resultCh := make(chan *domain.UsersIds)
	errCh := make(chan error)

	go func() {
		defer close(resultCh)
		defer close(errCh)

		err := validator.ValidatePercent(dto.Percent)
		if err != nil {
			errCh <- err
			return
		}
		userIds, err := s.storage.GetPercentOfUsersIds(ctx, dto)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- userIds
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

func (s userService) DeleteUsers(ctx context.Context, dto *domain.UsersIds) (*domain.UserAffected, error) {
	resultCh := make(chan *domain.UserAffected)
	errCh := make(chan error)

	go func() {
		defer close(resultCh)
		defer close(errCh)

		err := validator.ValidateIds(dto.Ids)
		if err != nil {
			errCh <- err
			return
		}
		userIds, err := s.storage.DeleteUsers(ctx, dto)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- userIds
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
