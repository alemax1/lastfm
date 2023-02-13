package repository

import (
	"database/sql"
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
	SaveTrackInfoToDB(domain.TracksResponse) error
}

func (t trackPG) SaveTrackInfoToDB(tracks domain.TracksResponse) (err error) {
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

			if e := txn.Commit(); e != nil {
				err = fmt.Errorf("commit: %v", err)
			}
		}
	}()

	for i := range tracks.Data {

		if re, err := txn.Exec("INSERT INTO artists(name) VALUES($1)", tracks.Data[i].Artist); err != nil {
			fmt.Println(re, err)
		}

		// if err := txn.QueryRow("INSERT INTO artists(name) VALUES($1) RETURNING id", tracks.Data[i].Artist).Scan(&artistID); err != nil {
		// 	return fmt.Errorf("exec artist: %v", err)
		// }

		// fmt.Println(artistID, "      3")

		// var trackID int

		// listeners, err := strconv.Atoi(tracks.Data[i].Listeners)
		// if err != nil {
		// 	return fmt.Errorf("strconv: %v", err)
		// }

		// playcount, err := strconv.Atoi(tracks.Data[i].Playcount)
		// if err != nil {
		// 	return fmt.Errorf("strconv: %v", err)
		// }

		// if err := txn.QueryRow("INSERT INTO tracks(name, listeners, playcounts, artist_id) VALUES($1, $2, $3, $4) RETURNING id",
		// 	tracks.Data[i].Name,
		// 	listeners,
		// 	playcount,
		// 	artistID,
		// ).Scan(&trackID); err != nil {
		// 	return fmt.Errorf("exec track: %v", err)
		// }

		// fmt.Println(trackID, "2")

		// for j := range tracks.Data[i].Tags {
		// 	var tagID int

		// 	if err := txn.QueryRow("INSERT INTO tags(name) VALUES($1) ON CONFLICT (name) DO NOTHING RETURNING(select id where name=$1)",
		// 		tracks.Data[i].Tags[j].Name,
		// 	).Scan(
		// 		&tagID,
		// 	); err != nil && !errors.Is(err, sql.ErrNoRows) {
		// 		return fmt.Errorf("exec tag: %v", err)
		// 	}

		// 	fmt.Println(tagID, "????????")

		// 	if tagID == 0 {
		// 		continue
		// 	}

		// 	if _, err := txn.Exec("INSERT INTO tracks_tags(track_id, tag_id) VALUES($1, $2)", trackID, tagID); err != nil {
		// 		return fmt.Errorf("exec tracksTags: %v", err)
		// 	}
		// }
	}

	return nil
}
