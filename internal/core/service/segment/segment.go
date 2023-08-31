package segment

import (
	"context"

	"segment-service/internal/core/domain"
	"segment-service/internal/core/ports/storage"
	"segment-service/internal/lib/validator"
)

type segmentService struct {
	storage storage.SegmentStorager
}

func NewSegmentService(storage storage.SegmentStorager) *segmentService {
	return &segmentService{
		storage: storage,
	}
}

func (s segmentService) AddSegment(ctx context.Context, dto *domain.SegmentAddDTO) (*domain.Segment, error) {
	resultCh := make(chan *domain.Segment)
	errCh := make(chan error)

	go func() {
		defer close(errCh)
		err := validator.ValidateSegmentName(dto.Name)
		if err != nil {
			errCh <- err
			return
		}
		result, err := s.storage.AddSegment(ctx, dto)
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
func (s segmentService) UpdateSegment(ctx context.Context, dto *domain.SegmentUpdateDTO) (*domain.Segment, error) {
	resultCh := make(chan *domain.Segment)
	errCh := make(chan error)

	go func() {
		defer close(errCh)
		err := validator.ValidateSegmentName(dto.NewName)
		if err != nil {
			errCh <- err
			return
		}
		result, err := s.storage.UpdateSegment(ctx, dto)
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

func (s segmentService) DeleteSegment(ctx context.Context, dto *domain.SegmentName) error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)
		err := validator.ValidateSegmentName(dto.Name)
		if err != nil {
			errCh <- err
			return
		}
		err = s.storage.DeleteSegment(ctx, dto)
		errCh <- err

	}()

	select {
	case <-ctx.Done():
		return domain.ErrorTimeOut
	case err := <-errCh:
		return err
	}
}

func (s segmentService) GetAllSegments(ctx context.Context) (*[]domain.Segment, error) {
	resultCh := make(chan *[]domain.Segment)
	errCh := make(chan error)

	go func() {
		defer close(errCh)

		result, err := s.storage.GetAllSegments(ctx)
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

func (s segmentService) GetSegmentByName(ctx context.Context, dto *domain.SegmentName) (*domain.Segment, error) {
	resultCh := make(chan *domain.Segment)
	errCh := make(chan error)

	go func() {
		defer close(resultCh)
		defer close(errCh)
		err := validator.ValidateSegmentName(dto.Name)
		if err != nil {
			errCh <- err
			return
		}
		result, err := s.storage.GetSegmentByName(ctx, dto)
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

func (s segmentService) CheckSegmentsExists(ctx context.Context, dto *domain.SegmentNames) error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)
		err := validator.ValidateSegmentNames(dto.Names)
		if err != nil {
			errCh <- err
			return
		}
		err = s.storage.CheckSegmentsExists(ctx, dto)
		errCh <- err

	}()

	select {
	case <-ctx.Done():
		return domain.ErrorTimeOut
	case err := <-errCh:
		return err
	}
}

func (s segmentService) GetSegmentsIds(ctx context.Context, dto *domain.SegmentNames) (*domain.SegmentIds, error) {
	resultCh := make(chan *domain.SegmentIds)
	errCh := make(chan error)

	go func() {
		defer close(resultCh)
		defer close(errCh)
		err := validator.ValidateSegmentNames(dto.Names)
		if err != nil {
			errCh <- err
			return
		}
		result, err := s.storage.GetSegmentsIds(ctx, dto)
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
