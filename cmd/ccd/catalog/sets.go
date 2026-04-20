package catalog

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/indexdata/ccms/internal/global"
)

func (c *Catalog) initSets() error {
	sql := "select setname from ccms.sets"
	rows, err := c.dp.Query(context.TODO(), sql)
	if err != nil {
		return fmt.Errorf("selecting sets: %v", err)
	}
	defer rows.Close()
	sets := make(map[string]struct{})
	for rows.Next() {
		var set string
		if err := rows.Scan(&set); err != nil {
			return fmt.Errorf("reading sets: %v", err)
		}
		sets[set] = struct{}{}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("reading sets: %v", err)
	}
	c.sets = sets
	return nil
}

func (c *Catalog) SetExists(setName string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if setName == "reserve" {
		return true
	}
	_, ok := c.sets[setName]
	return ok
}

func (c *Catalog) IsValidTargetSet(setName string) bool {
	if setName == "reserve" {
		return false
	}
	if !strings.ContainsRune(setName, '.') {
		return false
	}
	if strings.HasPrefix(setName, ".") || strings.HasSuffix(setName, ".") {
		return false
	}
	return c.ProjectExists(setName)
}

// return table containing set
func SetTable(setName string) string {
	if setName == "reserve" {
		return "ccms.reserve"
	}
	sp := strings.Split(setName, ".")
	return sp[0] + ".s_" + sp[1]
}

func (c *Catalog) AllSets() []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	sets := make([]string, len(c.sets)+1)
	i := 0
	sets[i] = "reserve"
	i++
	for k := range c.sets {
		sets[i] = k
		i++
	}
	sortSetNames(sets)
	return sets
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

func (c *Catalog) CreateSet(setName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	tx, err := c.dp.Begin(context.TODO())
	if err != nil {
		return fmt.Errorf("creating set %q: opening transaction: %v", setName, err)
	}
	defer tx.Rollback(context.TODO())

	sql := "create table " + SetTable(setName) + "(" +
		"id bigint primary key)"
	if _, err := tx.Exec(context.TODO(), sql); err != nil {
		return fmt.Errorf("creating set %q: %v", setName, err)
	}
	sql = "insert into ccms.sets (setname) values ($1)"
	if _, err := tx.Exec(context.TODO(), sql, setName); err != nil {
		return fmt.Errorf("registering set %q: %v", setName, global.PGErr(err))
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return fmt.Errorf("creating set %q: committing changes: %v", setName, err)
	}

	c.sets[setName] = struct{}{}
	return nil
}

func (c *Catalog) DropSet(setName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	tx, err := c.dp.Begin(context.TODO())
	if err != nil {
		return fmt.Errorf("dropping set %q: opening transaction: %v", setName, err)
	}
	defer tx.Rollback(context.TODO())

	sql := "drop table " + SetTable(setName)
	if _, err := tx.Exec(context.TODO(), sql); err != nil {
		return fmt.Errorf("dropping set %q: %v", setName, err)
	}
	sql = "delete from ccms.sets where setname=$1"
	if _, err := tx.Exec(context.TODO(), sql, setName); err != nil {
		return fmt.Errorf("deregistering set %q: %v", setName, global.PGErr(err))
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return fmt.Errorf("dropping set %q: committing changes: %v", setName, err)
	}

	delete(c.sets, setName)
	return nil
}
