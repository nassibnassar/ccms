package cat

import (
	"cmp"
	"errors"
	"fmt"
	"slices"

	"github.com/indexdata/ccms/internal/dbx"
	"github.com/indexdata/ccms/internal/pgerr"
	"github.com/indexdata/ccms/prop"
	"github.com/jackc/pgx/v5"
)

func CreateFund(d *dbx.DB, fund string) error {
	sql := "insert into ccms.fund (name, title) values ($1, $2)"
	if _, err := d.Q.Exec(d.C, sql, fund, makeTitle(fund)); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

// FundExists does not do synchronization and must not access the catalog cache
func FundExists(d *dbx.DB, fund string) (bool, error) {
	id, err := SelectFundID(d, fund)
	if err != nil {
		return false, err
	}
	return id != -1, nil
}

// returns fund id, or -1 if fund does not exist
func SelectFundID(d *dbx.DB, fund string) (int64, error) {
	var q = "select id from ccms.fund where name=$1"
	var id int64
	err := d.Q.QueryRow(d.C, q, fund).Scan(&id)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return -1, nil
	case err != nil:
		return 0, pgerr.Error(err)
	default:
		return id, nil
	}
}

func AllFunds(d *dbx.DB) ([]prop.Prop, error) {
	q := "select name, title from ccms.fund"
	rows, err := d.Q.Query(d.C, q)
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
	sortFundNames(funds)
	return funds, nil
}

func sortFundNames(funds []prop.Prop) {
	slices.SortFunc(funds, func(x, y prop.Prop) int {
		return cmp.Compare(x.Name, y.Name)
	})
}
