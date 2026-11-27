package cat

import (
	"errors"

	"github.com/indexdata/ccms/cmd/ccd/dberr"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
	"github.com/jackc/pgx/v5"
)

func LastHarvest(db *dbx.DB) (string, string, error) {
	sql := "select resumption_token, last_datestamp from ccms.harvest where name='main'"
	var resumptionToken, lastDatestamp string
	err := db.QueryRow(db.Ctx, sql).Scan(&resumptionToken, &lastDatestamp)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return "", "", errors.New("internal error: harvest state not found")
	case err != nil:
		return "", "", dberr.Error(err)
	default:
		return resumptionToken, lastDatestamp, nil
	}
}

func UpdateHarvest(db *dbx.DB, resumptionToken, lastDatestamp string) error {
	sql := "update ccms.harvest set resumption_token=$1, last_datestamp=$2 where name='main'"
	if _, err := db.Exec(db.Ctx, sql, resumptionToken, lastDatestamp); err != nil {
		return dberr.Error(err)
	}
	return nil
}

func RemoveResumptionToken(db *dbx.DB) error {
	sql := "update ccms.harvest set resumption_token='' where name='main'"
	if _, err := db.Exec(db.Ctx, sql); err != nil {
		return dberr.Error(err)
	}
	return nil
}
