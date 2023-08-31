package segment

import (
	"segment-service/internal/core/domain"
	mock_storage "segment-service/pkg/mocks/storage"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

type addOut struct {
	Result *domain.Segment
	Err    error
}

func TestAddSegment(t *testing.T) {
	cases := []struct {
		Name      string
		In        *domain.SegmentAddDTO
		Out       *addOut
		WhaiError bool
		Prepare   func(ctx context.Context, in *domain.SegmentAddDTO, storage mock_storage.MockSegmentStorager)
		TimeLinit time.Duration
	}{
		{
			Name: "no error",
			In:   &domain.SegmentAddDTO{Name: "SELL"},
			Out: &addOut{
				Result: &domain.Segment{Id: 1, Name: "SELL"},
				Err:    nil,
			},
			WhaiError: false,
			Prepare: func(ctx context.Context, in *domain.SegmentAddDTO, storage mock_storage.MockSegmentStorager) {
				storage.EXPECT().AddSegment(ctx, in).Return(&domain.Segment{Id: 1, Name: "SELL"}, nil)
			},
			TimeLinit: time.Second * 10,
		},
		{
			Name: "bad segment name error",
			In:   &domain.SegmentAddDTO{Name: ""},
			Out: &addOut{
				Result: nil,
				Err:    domain.ErrorInvalidName,
			},
			WhaiError: true,
			Prepare: nil,
			TimeLinit: time.Second * 10,
		},
		{
			Name: "time out error",
			In:   &domain.SegmentAddDTO{Name: ""},
			Out: &addOut{
				Result: nil,
				Err:    domain.ErrorTimeOut,
			},
			WhaiError: true,
			Prepare: nil,
			TimeLinit: 0,
		},
		{
			Name: "segment registered error",
			In:   &domain.SegmentAddDTO{Name: "SELL"},
			Out: &addOut{
				Result: nil,
				Err:    domain.ErrorSegmentExists,
			},
			WhaiError: true,
			Prepare: func(ctx context.Context, in *domain.SegmentAddDTO, storage mock_storage.MockSegmentStorager) {
				storage.EXPECT().AddSegment(ctx, in).Return(nil, domain.ErrorSegmentExists)
			},
			TimeLinit: time.Second * 10,
		},
	}

	for _, cc := range cases {
		t.Run(cc.Name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			ctx, cancel := context.WithTimeout(context.Background(), cc.TimeLinit)
			defer cancel()
			storage := mock_storage.NewMockSegmentStorager(ctl)
			service := NewSegmentService(storage)
			if cc.Prepare != nil {
				cc.Prepare(ctx, cc.In, *storage)
			}
			segment, err := service.AddSegment(ctx, cc.In)
			if err != nil && !cc.WhaiError{
				t.Errorf("don't want error, but it was given: %s", err.Error())
			}
			if cc.WhaiError {
				require.EqualError(t, err, cc.Out.Err.Error())
			} else {
				if segment.Id != cc.Out.Result.Id || segment.Name != cc.Out.Result.Name {
					t.Errorf("wanted - {'id': %d, 'name': %s}, have - {'id': %d, 'name': %s}", cc.Out.Result.Id, cc.Out.Result.Name, segment.Id, segment.Name)
				}
			}
		})
	}
}
