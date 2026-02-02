package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zmonsoonz/go-test-api/internal/content/domain"
	"github.com/zmonsoonz/go-test-api/internal/platform/db/postgres"
)

type TrackRep struct {
	db *sqlx.DB
}

func NewTrackRep(db *sqlx.DB) *TrackRep {
	return &TrackRep{db: db}
}
func (r *TrackRep) Create(track domain.Track, userId int) (int, error) {
	var id int
	libId, err := r.getLibraryId(userId)
	if err != nil {
		return 0, err
	}

	createTrackQuery := fmt.Sprintf("INSERT INTO %s (title, description, uploaded_at, library_id) VALUES ($1, $2, $3, $4) RETURNING id", postgres.TracksTable)
	row := r.db.QueryRow(createTrackQuery, track.Title, track.Description, track.UploadedAt, libId)
	err = row.Scan(&id);

	return id, err
}
func (r *TrackRep) GetAll(userId int) ([]domain.Track, error) {
	var tracks []domain.Track;

	query := fmt.Sprintf("SELECT tr.id, tr.library_id, tr.title, tr.description, tr.uploaded_at FROM %s tr INNER JOIN %s lb ON tr.library_id = lb.id WHERE lb.user_id = $1", 
	postgres.TracksTable, postgres.LibrariesTable)

	err := r.db.Select(&tracks, query, userId)	

	return tracks, err
}

func (r *TrackRep) GetById(id, userId int) (domain.Track, error) {
	var track domain.Track;

	query := fmt.Sprintf("SELECT tr.id, tr.description, tr.uploaded_at FROM %s tr INNER JOIN %s lb ON tr.library_id = lb.id WHERE lb.user_id = $1 AND tr.id = $2", 
	postgres.TracksTable, postgres.LibrariesTable)

	err := r.db.Get(&track, query, userId, id)

	return track, err
}

// func (r *TrackRep) Delete(id int) error {
// 	query := fmt.Sprintf("DELETE FROM %s tr INNER JOIN %s lb ON tr.library_id = lb.id WHERE lb.user_id = $1 AND tr.id = $2", 
// 	postgres.TracksTable, postgres.LibrariesTable)
// }

// func (r *TrackRep) Update()  error {

// }

func (r *TrackRep) getLibraryId(userID int) (int, error) {
	var id int

	query := `SELECT id FROM libraries WHERE user_id = $1`
	err := r.db.Get(&id, query, userID)

	return id, err
}