package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"segment-service/internal/core/domain"
	"segment-service/internal/core/ports/storage"
	"strconv"
)

type config interface {
	GetHistoryDir() string
	GetHttpURL() string
}

type historyService struct {
	storage       storage.HistoryStorage
	dir           string
	hileServerURL string
}

func NewHistoryService(s storage.HistoryStorage, cfg config) *historyService {
	return &historyService{
		storage:       s,
		dir:           cfg.GetHistoryDir(),
		hileServerURL: cfg.GetHttpURL(),
	}
}

func (s historyService) AddHistory(dto *domain.HistoryAddDTO) {
	s.storage.AddHistory(dto)
}

func (s historyService) GetHistoryOfUser(ctx context.Context, dto *domain.HistoryOfUserGetDTO) (*domain.HistoryResponce, error) {
	errCh := make(chan error)
	resultCh := make(chan *domain.HistoryResponce)

	go func() {
		defer close(resultCh)
		defer close(errCh)

		history, err := s.storage.GetHistoryOfUser(ctx, dto)
		if err != nil {
			errCh <- err
			return
		}
		if len(*history) == 0 {
			errCh <- domain.ErrorNoHistoryResult
			return
		}
		if _, err := os.Stat(s.dir); err != nil {
			err = os.Mkdir(s.dir, 0777)
			if err != nil {
				errCh <- err
				return
			}
		}
		filename := strconv.Itoa(int(dto.UserId)) + "." + strconv.Itoa(dto.Month) + "." + strconv.Itoa(dto.Year) + ".csv"
		path := fmt.Sprintf("./%s/%s", s.dir, filename)
		if _, err := os.Stat(path); err == nil {
			err := os.Remove(path)
			if err != nil {
				errCh <- err
				return
			}
		}
		file, err := os.Create(path)
		defer file.Close()

		if err != nil {
			errCh <- err
			return
		}
		header := []string{"segment_name", "action", "time"}
		writer := csv.NewWriter(file)
		defer writer.Flush()

		err = writer.Write(header)
		if err != nil {
			errCh <- err
			return
		}
		for _, record := range *history {
			err = writer.Write(record.Values())
			if err != nil {
				errCh <- err
				return
			}
		}

		resultCh <- &domain.HistoryResponce{URL: fmt.Sprintf("http://%s/%s/%s", s.hileServerURL, s.dir, filename)}
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
