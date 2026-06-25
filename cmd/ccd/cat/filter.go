package cat

import (
	"cmp"
	"errors"
	"slices"
	"strings"

	"github.com/indexdata/ccms/internal/dbx"
	"github.com/indexdata/ccms/internal/pgerr"
	"github.com/jackc/pgx/v5"
)

type Filter struct {
	Name       string
	Definition string
}

func FilterExists(d *dbx.DB, filter string) (bool, error) {
	sql := "select 1 from ccms.filter where name=$1"
	var n int32
	err := d.Q.QueryRow(d.C, sql, filter).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func FilterSQL(d *dbx.DB, filter string) (string, error) {
	rows, _ := d.Q.Query(d.C, "select sql from ccms.filter where name=$1", filter)
	filterSQL, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		return "", err
	}
	if len(filterSQL) == 0 {
		return "", errors.New("filter \"" + filter + "\" does not exist")
	}
	return filterSQL[0], nil
}

func Filters(d *dbx.DB) ([]Filter, error) {
	sql := "select name, command from ccms.filter"
	rows, err := d.Q.Query(d.C, sql)
	if err != nil {
		return nil, pgerr.Error(err)
	}
	defer rows.Close()
	filters := make([]Filter, 0)
	for rows.Next() {
		var name, command string
		if err := rows.Scan(&name, &command); err != nil {
			return nil, pgerr.Error(err)
		}
		filters = append(filters, Filter{
			Name:       name,
			Definition: command,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, pgerr.Error(err)
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
