package cat

import (
	"errors"
	"strings"

	"github.com/indexdata/ccms/cmd/ccd/dberr"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
	"github.com/indexdata/ccms/prop"
	"github.com/jackc/pgx/v5"
)

func CreateFund(db *dbx.DB, fund string) error {
	sql := "insert into ccms.fund (name, title) values ($1, $2)"
	if _, err := db.Exec(db.Ctx, sql, fund, makeTitle(fund)); err != nil {
		return dberr.Error(err)
	}
	return nil
}

// returns fund ID, or 0 if fund does not exist
func FundID(db *dbx.DB, fund string) (int32, error) {
	var q = "select id from ccms.fund where name=$1"
	var id int32
	err := db.QueryRow(db.Ctx, q, fund).Scan(&id)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return 0, nil
	case err != nil:
		return 0, dberr.Error(err)
	default:
		return id, nil
	}
}

func Funds(db *dbx.DB) (prop.Property, error) {
	sql := "select name, title from ccms.fund"
	rows, err := db.Query(db.Ctx, sql)
	if err != nil {
		return nil, dberr.Error(err)
	}
	funds, err := pgx.CollectRows(rows, pgx.RowToStructByPos[prop.Prop])
	if err != nil {
		return nil, err
	}
	return funds, nil
}

func IsValidFundName(fund string) bool {
	if strings.ContainsRune(fund, '.') {
		return false
	}
	return true
}
