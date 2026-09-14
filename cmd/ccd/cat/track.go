package cat

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/indexdata/ccms/cmd/ccd/dberr"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
	"github.com/indexdata/ccms/prop"
	"github.com/jackc/pgx/v5"
)

func CreateTrack(db *dbx.DB, track string) error {
	sql := "insert into ccms.track (name, title) values ($1, $2)"
	if _, err := db.Exec(db.Ctx, sql, track, makeTitle(track)); err != nil {
		return dberr.Error(err)
	}
	return nil
}

func DropTrack(db *dbx.DB, track string) error {
	trackID, err := TrackID(db, track)
	if err != nil {
		return err
	}
	if trackID == 0 {
		return errors.New("track \"" + track + "\" does not exist")
	}

	projects, err := ProjectsHavingTrack(db, trackID)
	if err != nil {
		return err
	}
	if len(projects) != 0 {
		slices.Sort(projects)
		for i := range projects {
			projects[i] = "\"" + projects[i] + "\""
		}
		var s string
		if len(projects) > 1 {
			s = "s"
		}
		return errors.New("track \"" + track + "\" is used in project" + s + " " + strings.Join(projects, ", "))
	}

	sql := "delete from ccms.track where id=$1"
	if _, err := db.Exec(db.Ctx, sql, trackID); err != nil {
		return dberr.Error(err)
	}
	return nil
}

// returns track ID, or 0 if track does not exist
func TrackID(db *dbx.DB, track string) (int32, error) {
	sql := "select id from ccms.track where name=$1"
	var id int32
	err := db.QueryRow(db.Ctx, sql, track).Scan(&id)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return 0, nil
	case err != nil:
		return 0, dberr.Error(err)
	default:
		return id, nil
	}
}

func Tracks(db *dbx.DB) (prop.Property, error) {
	sql := "select name, title from ccms.track"
	rows, _ := db.Query(db.Ctx, sql)
	tracks, err := pgx.CollectRows(rows, pgx.RowToStructByPos[prop.Prop])
	if err != nil {
		return nil, err
	}
	return tracks, nil
}

func IsValidTrackName(track string) bool {
	if strings.ContainsRune(track, '.') {
		return false
	}
	return true
}

func TrackProperties(db *dbx.DB, track string) ([][2]string, error) {
	var title string
	sql := `select f.title from ccms.track f where f.name=$1`
	err := db.QueryRow(db.Ctx, sql, track).Scan(&title)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, fmt.Errorf("track %q does not exist", track)
	case err != nil:
		return nil, dberr.Error(err)
	default:
	}
	prop := [][2]string{
		{"name", track},
		{"title", title},
	}
	return prop, nil
}

func AlterTrackSetProperty(db *dbx.DB, track, property, value string, stringLiteral bool) error {
	switch property {
	case "name":
		if stringLiteral {
			return invalidValueError(property, "'"+value+"'")
		}
		if value == "" {
			return invalidValueError(property, value)
		}
	case "title":
		if !stringLiteral {
			return invalidValueError(property, value)
		}
	default:
		return errors.New("property \"" + property + "\" does not exist")
	}

	sql := "update ccms.track set \"" + property + "\"=$1 where name=$2"
	if _, err := db.Exec(db.Ctx, sql, value, track); err != nil {
		return dberr.Error(err)
	}
	return nil
}
