package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"track/internal/domain"
)

type trackPG struct {
	Conn *sql.DB
}

func NewTrackPG(conn *sql.DB) TrackPG {
	return &trackPG{
		Conn: conn,
	}
}

type TrackPG interface {
	SaveTrackInfoToDB(ctx context.Context, tracks domain.TracksResponse) error
	GetTracksByTag(ctx context.Context, page, limit int, tag string) ([]domain.TrackDBResponse, error)
	GetTracksByArtist(ctx context.Context, page, limit int, artist string) ([]domain.TrackDBResponse, error)
}

func (t trackPG) SaveTrackInfoToDB(ctx context.Context, tracks domain.TracksResponse) (err error) {
	txn, err := t.Conn.Begin()
	if err != nil {
		return fmt.Errorf("begin: %v", err)
	}

	defer func() {
		if err != nil {
			if e := txn.Rollback(); e != nil {
				err = fmt.Errorf("%v: rollback: %v", err, e)

				return
			}
		}

		if e := txn.Commit(); e != nil {
			err = fmt.Errorf("commit: %v", err)
		}
	}()

	for i := range tracks.Data {
		var artistID int

		if err := txn.QueryRowContext(ctx,
			"INSERT INTO artists(name) VALUES($1) ON CONFLICT(name) DO NOTHING RETURNING id",
			tracks.Data[i].Artist).Scan(&artistID); err != nil && err != sql.ErrNoRows {
			return fmt.Errorf("exec artist: %v", err)
		}

		if artistID == 0 {
			if err := txn.QueryRowContext(ctx,
				"SELECT id from artists WHERE name = $1",
				tracks.Data[i].Artist).Scan(&artistID); err != nil {
				return fmt.Errorf("scan name: %v", err)
			}
		}

		var trackID int

		if err := txn.QueryRowContext(ctx,
			`INSERT INTO 
		tracks(name, listeners, playcounts, artist_id) 
			VALUES($1, $2, $3, $4) 
		ON CONFLICT (name, artist_id) DO NOTHING 
			RETURNING id`,
			tracks.Data[i].Name,
			tracks.Data[i].Listeners,
			tracks.Data[i].Playcount,
			artistID,
		).Scan(
			&trackID,
		); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("exec track: %v", err)
		}

		if trackID == 0 {
			if err := txn.QueryRowContext(ctx,
				"SELECT id from tracks WHERE name = $1",
				tracks.Data[i].Name).Scan(&trackID); err != nil {
				return fmt.Errorf("scan name: %v", err)
			}
		}

		for j := range tracks.Data[i].Tags {
			var tagID int

			if err := txn.QueryRowContext(ctx,
				"INSERT INTO tags(name) VALUES($1) ON CONFLICT (name) DO NOTHING RETURNING(id)",
				tracks.Data[i].Tags[j].Name).Scan(&tagID); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("exec tag: %v", err)
			}

			if tagID == 0 {
				if err := txn.QueryRowContext(ctx,
					"SELECT id from tags WHERE name = $1",
					tracks.Data[i].Tags[j].Name).Scan(&tagID); err != nil {
					return fmt.Errorf("scan name: %v", err)
				}
			}

			if _, err := txn.ExecContext(ctx,
				"INSERT INTO tracks_tags(track_id, tag_id) VALUES($1, $2) ON CONFLICT(track_id, tag_id) DO NOTHING",
				trackID, tagID); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("exec tracksTags: %v", err)
			}
		}
	}

	return nil
}

func (t trackPG) GetTracksByTag(ctx context.Context, page, limit int, tag string) ([]domain.TrackDBResponse, error) {
	rows, err := t.Conn.QueryContext(ctx,
		`SELECT 
	t.id, t.name, t.listeners, t.playcounts
		FROM tracks t
	JOIN artists a ON t.artist_id = a.id
	JOIN tracks_tags tt ON t.id = tt.track_id
	JOIN tags tg ON tt.tag_id = tg.id
	WHERE tg.name = $1
		ORDER BY t.id ASC
	LIMIT $2
	OFFSET $3`, tag, limit, page*limit)
	if err != nil {
		return nil, fmt.Errorf("query: %v", err)
	}

	var tracks []domain.TrackDBResponse

	for rows.Next() {
		var track domain.TrackDBResponse

		if err := rows.Scan(&track.ID, &track.Name, &track.Listeners, &track.Playcounts); err != nil {
			return nil, fmt.Errorf("scan: %v", err)
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (t trackPG) GetTracksByArtist(ctx context.Context, page, limit int, artist string) ([]domain.TrackDBResponse, error) {
	rows, err := t.Conn.QueryContext(ctx,
		`SELECT 
	t.id, t.name, t.listeners, t.playcounts
		FROM tracks t 
	JOIN artists a ON t.artist_id = a.id 
		WHERE a.name = $1 
	ORDER BY t.id 
	LIMIT $2 
	OFFSET $3`, artist, limit, page*limit)
	if err != nil {
		return nil, fmt.Errorf("query: %v", err)
	}

	var tracks []domain.TrackDBResponse

	for rows.Next() {
		var track domain.TrackDBResponse

		if err := rows.Scan(&track.ID, &track.Name, &track.Listeners, &track.Playcounts); err != nil {
			return nil, fmt.Errorf("scan: %v", err)
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
}
