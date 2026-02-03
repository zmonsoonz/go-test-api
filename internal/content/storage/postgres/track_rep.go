package repository

import (
	"errors"
	"fmt"
	"strings"

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

	query := fmt.Sprintf("SELECT tr.id,  tr.library_id, tr.title, tr.description, tr.uploaded_at FROM %s tr INNER JOIN %s lb ON tr.library_id = lb.id WHERE lb.user_id = $1 AND tr.id = $2", 
	postgres.TracksTable, postgres.LibrariesTable)

	err := r.db.Get(&track, query, userId, id)

	return track, err
}

func (r *TrackRep) Delete(id, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s tr USING %s lb WHERE tr.library_id = lb.id AND lb.user_id = $1 AND tr.id = $2", 
	postgres.TracksTable, postgres.LibrariesTable)

	_, err := r.db.Exec(query, userId, id)

	return err
}

func (r *TrackRep) Update(input domain.UpdateTrackInput, id, userId int)  error {
	setParts := []string{}
	args := []interface{}{}
	argId := 1

	if input.Title != nil {
		setParts = append(setParts, fmt.Sprintf("title = $%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setParts = append(setParts, fmt.Sprintf("description = $%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if len(setParts) == 0 {
		return errors.New("no fields to update")
	}

	query := fmt.Sprintf(`
		UPDATE tracks tr
		SET %s
		FROM libraries lb
		WHERE tr.library_id = lb.id
		  AND lb.user_id = $%d
		  AND tr.id = $%d
	`,
		strings.Join(setParts, ", "),
		argId,
		argId+1,
	)

	args = append(args, userId, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TrackRep) getLibraryId(userID int) (int, error) {
	var id int

	query := `SELECT id FROM libraries WHERE user_id = $1`
	err := r.db.Get(&id, query, userID)

	return id, err
}