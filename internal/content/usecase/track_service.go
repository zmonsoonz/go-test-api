package usecase

import (
	"time"

	"github.com/zmonsoonz/go-test-api/internal/content/domain"
	"github.com/zmonsoonz/go-test-api/internal/content/ports"
)

type TrackService struct {
	repository ports.TrackRepository
}

func NewTrackService(repository ports.TrackRepository) *TrackService {
	return &TrackService{repository: repository}
}
func (s *TrackService) Create(track domain.Track, userId int) (int, error) {
	track.UploadedAt = time.Now().Unix()
	return s.repository.Create(track, userId)
}

func (s *TrackService) GetAll(userId int) ([]domain.Track, error) {
	return s.repository.GetAll(userId)
}

func (s *TrackService) GetById(id int, userId int) (domain.Track, error) {
	return s.repository.GetById(id, userId)
}
