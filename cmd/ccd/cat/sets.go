package cat

import (
	"cmp"
	"errors"
	"slices"
	"strings"

	"github.com/indexdata/ccms/internal/dbx"
	"github.com/indexdata/ccms/internal/pgerr"
	"github.com/indexdata/ccms/internal/set"
	"github.com/jackc/pgx/v5"
)

func SetExists(d *dbx.DB, set set.Set) (bool, error) {
	if set.Set == "object" {
		projectExists, err := ProjectExists(d, set.Project)
		if err != nil {
			return false, err
		}
		return projectExists, nil
	}

	var q = "select 1 from ccms.sets where project=$1 and set=$2"
	var n int32
	err := d.Q.QueryRow(d.C, q, set.Project, set.Set).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func IsValidTargetSet(d *dbx.DB, set set.Set) (bool, error) {
	if set.Project == "" || set.Set == "" {
		return false, nil
	}
	if set.Set == "object" {
		return false, nil
	}
	projectExists, err := ProjectExists(d, set.Project)
	if err != nil {
		return false, err
	}
	return projectExists, nil
}

// return table containing set
func SetTable(set set.Set) string {
	if set.Set == "object" {
		return set.String()
	}
	return set.Project + ".s_" + set.Set
}

func AllSets(d *dbx.DB) ([]string, error) {
	sql := "select project||'.'||set from ccms.sets"
	rows, _ := d.Q.Query(d.C, sql)
	sets, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		return nil, pgerr.Error(err)
	}
	slices.Sort(sets)
	return sets, nil
}

func SetsInProject(d *dbx.DB, project string) ([]string, error) {
	sql := "select project||'.'||set from ccms.sets where project=$1"
	rows, _ := d.Q.Query(d.C, sql, project)
	sets, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		return nil, pgerr.Error(err)
	}
	slices.Sort(sets)
	return sets, nil
}

func sortSetNames(sets []string) {
	slices.SortFunc(sets, func(x, y string) int {
		a := !strings.ContainsRune(x, '.')
		b := !strings.ContainsRune(y, '.')
		if a && !b {
			return -1
		}
		if !a && b {
			return 1
		}
		return cmp.Compare(x, y)
	})
}

func CreateSet(d *dbx.DB, set set.Set) error {
	sql := "create table " + SetTable(set) + "(" +
		"id bigint primary key)"
	if _, err := d.Q.Exec(d.C, sql); err != nil {
		return pgerr.Error(err)
	}
	sql = "insert into ccms.sets (project, set) values ($1, $2)"
	if _, err := d.Q.Exec(d.C, sql, set.Project, set.Set); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func DropSet(d *dbx.DB, set set.Set) error {
	q := "drop table " + SetTable(set)
	if _, err := d.Q.Exec(d.C, q); err != nil {
		return pgerr.Error(err)
	}
	q = "delete from ccms.sets where project=$1 and set=$2"
	if _, err := d.Q.Exec(d.C, q, set.Project, set.Set); err != nil {
		return pgerr.Error(err)
	}
	return nil
}
