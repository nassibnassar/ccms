package cat

import (
	"errors"
	"fmt"
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
func FundID(db *dbx.DB, fund string) (int64, error) {
	var q = "select id from ccms.fund where name=$1"
	var id int64
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
	q := "select name, title from ccms.fund"
	rows, err := db.Query(db.Ctx, q)
	if err != nil {
		return nil, fmt.Errorf("selecting funds: %v", err)
	}
	defer rows.Close()
	funds := make([]prop.Prop, 0)
	for rows.Next() {
		var f prop.Prop
		if err := rows.Scan(&f.Name, &f.Title); err != nil {
			return nil, fmt.Errorf("reading funds: %v", err)
		}
		funds = append(funds, f)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("reading funds: %v", err)
	}
	return funds, nil
}

func IsValidFundName(fund string) bool {
	if strings.ContainsRune(fund, '.') {
		return false
	}
	return true
}
