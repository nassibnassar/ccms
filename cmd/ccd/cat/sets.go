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

func SetExists(d *dbx.DB, set string) (bool, error) {
	s := strings.Split(set, ".")
	if len(s) != 2 {
		return false, nil
	}
	if s[1] == "object" {
		projectExists, err := ProjectExists(d, s[0])
		if err != nil {
			return false, err
		}
		return projectExists, nil
	}

	var q = "select 1 from ccms.sets where setname=$1"
	var n int32
	err := d.Q.QueryRow(d.C, q, set).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func IsValidTargetSet(d *dbx.DB, set string) (bool, error) {
	s := strings.Split(set, ".")
	if len(s) != 2 {
		return false, nil
	}
	if s[0] == "" || s[1] == "" {
		return false, nil
	}
	if s[1] == "object" {
		return false, nil
	}
	projectExists, err := ProjectExists(d, set)
	if err != nil {
		return false, err
	}
	return projectExists, nil
}

// return table containing set
func SetTable(set string) string {
	s := strings.Split(set, ".")
	if s[1] == "object" {
		return set
	}
	return s[0] + ".s_" + s[1]
}

func SplitSchemaTable(schemaTable string) (string, string) {
	s := strings.Split(schemaTable, ".")
	switch len(s) {
	case 1:
		return s[0], ""
	case 2:
		return s[0], s[1]
	default:
		return "", ""
	}
}

func AllSets(d *dbx.DB) ([]string, error) {
	sql := "select setname from ccms.sets"
	rows, _ := d.Q.Query(d.C, sql)
	sets, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		return nil, pgerr.Error(err)
	}
	slices.Sort(sets)
	return sets, nil
}

func SetsInProject(d *dbx.DB, project string) ([]string, error) {
	sql := "select setname from ccms.sets where setname like $1"
	rows, _ := d.Q.Query(d.C, sql, project+".%")
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

func CreateSet(d *dbx.DB, set string) error {
	sql := "create table " + SetTable(set) + "(" +
		"id bigint primary key)"
	if _, err := d.Q.Exec(d.C, sql); err != nil {
		return pgerr.Error(err)
	}
	sql = "insert into ccms.sets (setname) values ($1)"
	if _, err := d.Q.Exec(d.C, sql, set); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func DropSet(d *dbx.DB, setName string) error {
	q := "drop table " + SetTable(setName)
	if _, err := d.Q.Exec(d.C, q); err != nil {
		return pgerr.Error(err)
	}
	q = "delete from ccms.sets where setname=$1"
	if _, err := d.Q.Exec(d.C, q, setName); err != nil {
		return pgerr.Error(err)
	}
	return nil
}
