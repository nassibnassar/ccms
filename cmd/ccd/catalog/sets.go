package catalog

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/indexdata/ccms/internal/pgerr"
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

func (c *Catalog) SetExists(set string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	s := strings.Split(set, ".")
	if len(s) != 2 {
		return false
	}
	if s[1] == "object" {
		return c.projectExists(s[0])
	}

	_, ok := c.sets[set]
	return ok
}

// IsValidTargetSet does not do synchronization and must not access the catalog cache
func (c *Catalog) IsValidTargetSet(set string) bool {
	s := strings.Split(set, ".")
	if len(s) != 2 {
		return false
	}
	if s[0] == "" || s[1] == "" {
		return false
	}
	if s[1] == "object" {
		return false
	}
	return c.ProjectExists(set)
}

// return table containing set
// SetTable does not do synchronization and must not access the catalog cache
func SetTable(set string) string {
	s := strings.Split(set, ".")
	if s[1] == "object" {
		return set
	}
	return s[0] + ".s_" + s[1]
}

func (c *Catalog) AllSets() []string {
	c.mu.Lock()
	sets := slices.Collect(maps.Keys(c.sets))
	c.mu.Unlock()

	sortSetNames(sets)
	return sets
}

func (c *Catalog) SetsInProject(project string) []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.setsInProject(project)
}

func (c *Catalog) setsInProject(project string) []string {
	prefix := project + "."
	sets := make([]string, 0)
	for k := range c.sets {
		if strings.HasPrefix(k, prefix) {
			sets = append(sets, k)
		}
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
		return errors.New("registering set \"" + setName + "\": " + pgerr.String(err))
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
		return errors.New("deregistering set \"" + setName + "\": " + pgerr.String(err))
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return fmt.Errorf("dropping set %q: committing changes: %v", setName, err)
	}

	delete(c.sets, setName)
	return nil
}
