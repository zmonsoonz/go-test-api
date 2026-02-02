package ports

import (
	"github.com/zmonsoonz/go-test-api/internal/content/domain"
)

type TrackService interface {
	GetAll(userId int) ([]domain.Track, error)
	Create(track domain.Track, userId int) (int, error)
	GetById(id int, userId int) (domain.Track, error)
	// Update(c gin.Context, track domain.Track) error
	// Delete(c gin.Context, id int) error
}

type TrackRepository interface {
	GetAll(userId int) ([]domain.Track, error)
	Create(track domain.Track, userId int) (int, error)
	GetById(id int, userId int) (domain.Track, error)
	// Update(c gin.Context, track domain.Track) error
	// Delete(c gin.Context, id int) error
}