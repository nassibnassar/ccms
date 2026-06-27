package cat

import (
	"cmp"
	"errors"
	"slices"
	"strings"

	"github.com/indexdata/ccms/cmd/ccd/dberr"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
	"github.com/jackc/pgx/v5"
)

type Filter struct {
	Name       string
	Definition string
}

func FilterExists(db *dbx.DB, filter string) (bool, error) {
	sql := "select 1 from ccms.filter where name=$1"
	var n int32
	err := db.QueryRow(db.Ctx, sql, filter).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, dberr.Error(err)
	default:
		return true, nil
	}
}

func FilterSQL(db *dbx.DB, filter string) (string, error) {
	rows, err := db.Query(db.Ctx, "select sql from ccms.filter where name=$1", filter)
	if err != nil {
		return "", dberr.Error(err)
	}
	filterSQL, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		return "", err
	}
	if len(filterSQL) == 0 {
		return "", errors.New("filter \"" + filter + "\" does not exist")
	}
	return filterSQL[0], nil
}

func Filters(db *dbx.DB) ([]Filter, error) {
	sql := "select name, command from ccms.filter"
	rows, err := db.Query(db.Ctx, sql)
	if err != nil {
		return nil, dberr.Error(err)
	}
	defer rows.Close()
	filters := make([]Filter, 0)
	for rows.Next() {
		var name, command string
		if err := rows.Scan(&name, &command); err != nil {
			return nil, dberr.Error(err)
		}
		filters = append(filters, Filter{
			Name:       name,
			Definition: command,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, dberr.Error(err)
	}
	return filters, nil
}

func SortFilters(filters []Filter) {
	slices.SortFunc(filters, func(x, y Filter) int {
		return cmp.Compare(x.Name, y.Name)
	})
}

func IsValidFilterName(filter string) bool {
	if strings.ContainsRune(filter, '.') {
		return false
	}
	return true
}
