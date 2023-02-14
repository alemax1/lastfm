package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"track/internal/domain"
)

type albumPG struct {
	conn *sql.DB
}

func NewAlbumPG(conn *sql.DB) AlbumPG {
	return &albumPG{
		conn: conn,
	}
}

type AlbumPG interface {
	SaveAlbumInfoToDB(domain.AlbumResponse) error
}

func (a albumPG) SaveAlbumInfoToDB(albumResp domain.AlbumResponse) (err error) {
	if len(albumResp.Tracks) == 0 {
		return
	}

	txn, err := a.conn.Begin()
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

	var artistID int

	if err := txn.QueryRow("INSERT INTO artists(name) VALUES($1) ON CONFLICT(name) DO NOTHING RETURNING id",
		albumResp.Artist).Scan(
		&artistID,
	); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("insert artists: %v", err)
	}

	if artistID == 0 {
		if err := txn.QueryRow("SELECT id FROM artists where name = $1", albumResp.Artist).Scan(&artistID); err != nil {
			return fmt.Errorf("select artistsID: %v", err)
		}
	}

	var albumID int

	if err := txn.QueryRow(`INSERT INTO 
		albums(name, listeners, playcounts, artist_id) 
	VALUES($1, $2, $3, $4) 
		ON CONFLICT (name, artist_id) DO NOTHING 
	RETURNING id`,
		albumResp.Name,
		albumResp.Listeners,
		albumResp.Playcount,
		artistID).Scan(
		&albumID,
	); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("isert albums: %v", err)
	}

	if albumID == 0 {
		if err := txn.QueryRow("SELECT id FROM albums WHERE name = $1", albumResp.Name).Scan(&albumID); err != nil {
			return fmt.Errorf("select albumID: %v", err)
		}
	}

	for i := range albumResp.Tags.Tags {
		var tagID int

		if err := txn.QueryRow("INSERT INTO tags(name) VALUES($1) ON CONFLICT(name) DO NOTHING RETURNING id", albumResp.Tags.Tags[i].Name).Scan(&tagID); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("insert tags: %v", err)
		}

		if tagID == 0 {
			if err := txn.QueryRow("SELECT id FROM tags WHERE name = $1", albumResp.Tags.Tags[i].Name).Scan(&tagID); err != nil {
				return fmt.Errorf("select tagsID: %v", err)
			}
		}

		if _, err := txn.Exec("INSERT INTO albums_tags(album_id, tag_id) VALUES($1, $2) ON CONFLICT (album_id, tag_id) DO NOTHING", albumID, tagID); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("insert a_tags: %v", err)
		}
	}

	for i := range albumResp.Tracks {
		var trackID int

		if err := txn.QueryRow("INSERT INTO tracks(name, listeners, playcounts, artist_id) VALUES($1, $2, $3, $4) ON CONFLICT (name, artist_id) DO NOTHING RETURNING id",
			albumResp.Tracks[i].Name,
			albumResp.Tracks[i].Listeners,
			albumResp.Tracks[i].Playcount,
			artistID,
		).Scan(&trackID); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("insert tracks: %v", err)
		}

		if trackID == 0 {
			if err := txn.QueryRow("SELECT id FROM tracks WHERE name = $1", albumResp.Tracks[i].Name).Scan(&trackID); err != nil {
				return fmt.Errorf("select trackID: %v", err)
			}
		}

		if _, err := txn.Exec("INSERT INTO albums_tracks(album_id, track_id) VALUES($1, $2) ON CONFLICT(album_id, track_id) DO NOTHING", albumID, trackID); err != nil {
			return fmt.Errorf("insert aTracks: %v", err)
		}

		for j := range albumResp.Tracks[i].Tags.Tags {
			var tagID int

			if err := txn.QueryRow("INSERT INTO tags(name) VALUES($1) ON CONFLICT (name) DO NOTHING RETURNING id", albumResp.Tracks[i].Tags.Tags[j].Name).Scan(&tagID); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("insert tags: %v", err)
			}

			if tagID == 0 {
				if err := txn.QueryRow("SELECT id FROM tags WHERE name = $1", albumResp.Tracks[i].Tags.Tags[j].Name).Scan(&tagID); err != nil {
					return fmt.Errorf("select tagsID: %v", err)
				}
			}

			if _, err := txn.Exec("INSERT INTO tracks_tags(track_id, tag_id) VALUES($1, $2) ON CONFLICT (track_id, tag_id) DO NOTHING", trackID, tagID); err != nil {
				return fmt.Errorf("insert tTags: %v", err)
			}
		}
	}

	return nil
}
