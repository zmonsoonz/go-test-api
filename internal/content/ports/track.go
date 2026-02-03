package ports

import (
	"github.com/zmonsoonz/go-test-api/internal/content/domain"
)

type TrackService interface {
	GetAll(userId int) ([]domain.Track, error)
	Create(track domain.Track, userId int) (int, error)
	GetById(id int, userId int) (domain.Track, error)
	Update(input domain.UpdateTrackInput, id, userId int) error
	Delete(id int, userId int) error
}

type TrackRepository interface {
	GetAll(userId int) ([]domain.Track, error)
	Create(track domain.Track, userId int) (int, error)
	GetById(id int, userId int) (domain.Track, error)
	Update(input domain.UpdateTrackInput, id, userId int) error
	Delete(id, userId int) error
}